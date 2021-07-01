package state

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
)

type State interface {
	PersistOrders(orders []*gen.Order) error 
	
	GetPreparedOrders() []*gen.PreparedOrder // TODO: you don;t need this, doesn't make sense for the state to know what is a prepared order

	GetItemInfo(item *gen.Item) (ItemInfo, error)
	MarkItemInCubby(orderId string, itemIndex int) error

	GetRemainingItemsCount() int

	GetOrderInfo(orderId string) (OrderInfo, error)
	GetAllOrderIds() ([]string, error) // TODO: make this return the whole state

	// TODO: these three methods should belong to the newly created order processor
	IsCurrentlyProcessing() bool
	SetCurrentlyProccessingFalse()
	SetCurrentlyProccessingTrue()
}

func New() State {
	return &state{
		// TODO: in the current state there can be one map orderid -> order info (where order indo contains info about the items)
		orderIdToItems:   make(map[string][]*itemStatus),
		orderIdToCubbyId: make(map[string]string),
	}
}

type state struct {
	mu                  sync.RWMutex // TODO: change mutex
	orderIdToItems      map[string][]*itemStatus
	orderIdToCubbyId    map[string]string
	currentlyProcessing bool
}

func (s *state) PersistOrders(orders []*gen.Order) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(orders) > 10 {
		return errors.New("only up to 10 orders allowed for now")
	}
	if len(s.orderIdToCubbyId) >= 10 {
		return errors.New("order cache full")
	}

	for _, order := range orders {
		cubbyIdForOrder := s.determineCubbyIdForOrder(order)

		// TODO: this will get simpler if we have the map to order info
		s.orderIdToCubbyId[order.Id] = cubbyIdForOrder
		fmt.Printf("order: %v -> cubby: %v\n", order.Id, cubbyIdForOrder)

		s.orderIdToItems[order.Id] = make([]*itemStatus, len(order.Items))

		for _, item := range order.Items {
			s.orderIdToItems[order.Id] = append(s.orderIdToItems[order.Id], &itemStatus{item: item, inCubby: false})
		}
	}

	return nil
}

// TODO: useless
func (s *state) GetPreparedOrders() []*gen.PreparedOrder {
	s.mu.Lock()
	defer s.mu.Unlock()

	preparedOrders := make([]*gen.PreparedOrder, len(s.orderIdToItems))
	for orderId, itemsWithStatus := range s.orderIdToItems {
		items := []*gen.Item{}

		for _, itemStatus := range itemsWithStatus {
			items = append(items, itemStatus.item)
		}

		preparedOrders = append(preparedOrders, &gen.PreparedOrder{
			Order: &gen.Order{
				Id:    orderId,
				Items: items,
			},
			Cubby: &gen.Cubby{
				Id: s.orderIdToCubbyId[orderId],
			},
		})
	}

	return preparedOrders
}

func (s *state) GetItemInfo(item *gen.Item) (ItemInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for orderId, items := range s.orderIdToItems {
		for i, itemStatus := range items {
			if itemStatus == nil || itemStatus.inCubby {
				continue
			}

			itemsEqual, err := itemsEqual(item, itemStatus.item)
			if err != nil {
				return ItemInfo{}, errors.New(err.Error())
			}

			if itemsEqual {
				return ItemInfo{
					OrderId: orderId,
					Index:   i,
					CubbyId: s.orderIdToCubbyId[orderId],
				}, nil
			}
		}
	}

	return ItemInfo{}, errors.New("there is no such item in any of the loaded orders")
}

func (s *state) MarkItemInCubby(orderId string, itemIndex int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	items, ok := s.orderIdToItems[orderId]
	if !ok {
		return errors.New("no such order")
	}
	if len(items) <= itemIndex {
		return errors.New("no such item in order")
	}
	s.orderIdToItems[orderId][itemIndex].inCubby = true

	// TODO: you can mark an order as fulfilled if all items are in a cuby
	return nil
}

func (s *state) GetRemainingItemsCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	count := 0
	for _, itemsStatuses := range s.orderIdToItems {
		itemsNotInCubbies := 0

		for _, itemStatus := range itemsStatuses {
			if itemStatus == nil {
				continue
			}
			if !itemStatus.inCubby {
				itemsNotInCubbies++
			}
		}

		count += itemsNotInCubbies
	}
	return count
}

func (s *state) determineCubbyIdForOrder(order *gen.Order) string {
	cubbyId := ordertocubby.Map(order.Id, 1, 10)
	for times := 2; s.isCubbyOccupied(cubbyId); times++ {
		cubbyId = ordertocubby.Map(order.Id, uint32(times), 10)
	}

	return cubbyId
}

func (s *state) isCubbyOccupied(cubbyId string) bool {
	occupiedCubbies := []string{}
	for _, id := range s.orderIdToCubbyId {
		if id != "" {
			occupiedCubbies = append(occupiedCubbies, id)
		}
	}
	for _, oc := range occupiedCubbies {
		if oc == cubbyId {
			return true
		}
	}

	return false
}

func (s *state) GetOrderInfo(orderId string) (OrderInfo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	itemStatuses, ok := s.orderIdToItems[orderId]
	if !ok {
		return OrderInfo{}, errors.New("no such order")
	}
	state := gen.OrderState_READY
	items := []*gen.Item{}
	for _, itemStatus := range itemStatuses {
		items = append(items, itemStatus.item)
		if !itemStatus.inCubby {
			state = gen.OrderState_PENDING
		}
	}
	cubbyId, ok := s.orderIdToCubbyId[orderId]
	if !ok {
		return OrderInfo{}, errors.New("no cubby for order")
	}

	return OrderInfo{
		State: state,
		Order: &gen.Order{
			Id:    orderId,
			Items: items,
		},
		Cubby: &gen.Cubby{
			Id: cubbyId,
		},
	}, nil
}

func (s *state) GetAllOrderIds() ([]string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ordersCount := len(s.orderIdToItems)
	if ordersCount == 0 {
		return nil, errors.New("there are currently no orders")
	}
	orderIds := make([]string, len(s.orderIdToItems))
	for orderId := range s.orderIdToItems {
		orderIds = append(orderIds, orderId)
	}

	return orderIds, nil
}

func (s *state) IsCurrentlyProcessing() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.currentlyProcessing
}

func (s *state) SetCurrentlyProccessingFalse() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentlyProcessing = false
}

func (s *state) SetCurrentlyProccessingTrue() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.currentlyProcessing = true
}
