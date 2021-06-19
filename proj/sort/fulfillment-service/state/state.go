package state

import (
	"errors"
	"fmt"
	"sync"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
)

type State interface {
	PersistOrders(in *gen.LoadOrdersRequest) error
	GetCompleteResponse() *gen.CompleteResponse
	GetItemInfo(item *gen.Item) (ItemInfoModel, error)
	RemoveItemFromOrder(orderId string, itemIndex int) error
	GetRemainingItemsCount() int
}

func NewState() State {
	return &state{
		orderIdItems:   make(map[string][]*gen.Item),
		orderIdCubbyId: make(map[string]string),
	}
}

type state struct {
	mu             sync.RWMutex
	orderIdItems   map[string][]*gen.Item
	orderIdCubbyId map[string]string
}

func (s *state) PersistOrders(in *gen.LoadOrdersRequest) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(in.Orders) > 10 {
		return errors.New("only up to 10 orders allowed for now")
	}

	if len(s.orderIdCubbyId) >= 10 {
		return errors.New("order cache full")
	}

	for _, order := range in.Orders {
		cubbyIdForOrder := s.determineCubbyIdForOrder(order)

		s.orderIdCubbyId[order.Id] = cubbyIdForOrder
		fmt.Printf("order: %v -> cubby: %v\n", order.Id, cubbyIdForOrder)

		s.orderIdItems[order.Id] = make([]*gen.Item, len(order.Items))
		copy(s.orderIdItems[order.Id], order.Items)
	}

	return nil
}

func (s *state) GetCompleteResponse() *gen.CompleteResponse {
	s.mu.RLock()
	defer s.mu.RUnlock()

	preparedOrders := make([]*gen.PreparedOrder, len(s.orderIdItems))

	for orderId, items := range s.orderIdItems {
		preparedOrders = append(preparedOrders, &gen.PreparedOrder{
			Order: &gen.Order{
				Id:    orderId,
				Items: items,
			},
			Cubby: &gen.Cubby{
				Id: s.orderIdCubbyId[orderId],
			},
		})
	}

	return &gen.CompleteResponse{
		Status: "ok?",
		Orders: preparedOrders,
	}
}

func (s *state) GetItemInfo(item *gen.Item) (ItemInfoModel, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for orderId, items := range s.orderIdItems {
		for i, itm := range items {
			itemsEqual, err := itemsEqual(item, itm)
			if err != nil {
				return ItemInfoModel{}, errors.New(err.Error())
			}
			if itemsEqual {
				return ItemInfoModel{
					OrderId: orderId,
					Index:   i,
					CubbyId: s.orderIdCubbyId[orderId],
				}, nil
			}
		}
	}

	return ItemInfoModel{}, errors.New("there is no such item in any of the loaded orders")
}

func (s *state) RemoveItemFromOrder(orderId string, itemIndex int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orderIdItems[orderId][itemIndex] = s.orderIdItems[orderId][len(s.orderIdItems[orderId])-1]
	s.orderIdItems[orderId] = s.orderIdItems[orderId][:len(s.orderIdItems[orderId])-1]
	return nil // think what could go wrong in this method
}

func (s *state) GetRemainingItemsCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	for _, items := range s.orderIdItems {
		count += len(items)
	}
	return count
}

func (s *state) determineCubbyIdForOrder(order *gen.Order) string {
	cubbyId := ordertocubby.Map(order.Id, 1, 10)
	for times := 2; s.cubbyIsOccupied(cubbyId); times++ {
		cubbyId = ordertocubby.Map(order.Id, uint32(times), 10)
	}

	return cubbyId
}

func (s *state) cubbyIsOccupied(cubby string) bool {
	occupiedCubbies := []string{}
	for _, cubby := range s.orderIdCubbyId {
		if cubby != "" {
			occupiedCubbies = append(occupiedCubbies, cubby)
		}
	}
	for _, oc := range occupiedCubbies {
		if oc == cubby {
			return true
		}
	}

	return false
}

func itemsEqual(i1 *gen.Item, i2 *gen.Item) (bool, error) {
	if i1 == nil || i2 == nil {
		return false, errors.New("cannot compare nil items")
	}

	return i1.Code == i2.Code && i1.Label == i2.Label, nil
}
