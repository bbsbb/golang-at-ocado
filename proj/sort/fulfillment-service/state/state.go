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
	GetPreparedOrders() []*gen.PreparedOrder
	GetItemInfo(item *gen.Item) (ItemInfoModel, error)
	RemoveItemFromOrder(orderId string, itemIndex int) error
	GetRemainingItemsCount() int
}

func New() State {
	return &state{
		orderIdToItems:   make(map[string][]*gen.Item),
		orderIdToCubbyId: make(map[string]string),
	}
}

type state struct {
	mu               sync.RWMutex
	orderIdToItems   map[string][]*gen.Item
	orderIdToCubbyId map[string]string
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

		s.orderIdToCubbyId[order.Id] = cubbyIdForOrder
		fmt.Printf("order: %v -> cubby: %v\n", order.Id, cubbyIdForOrder)

		s.orderIdToItems[order.Id] = make([]*gen.Item, len(order.Items))
		copy(s.orderIdToItems[order.Id], order.Items)
	}

	return nil
}

func (s *state) GetPreparedOrders() []*gen.PreparedOrder {
	s.mu.RLock()
	defer s.mu.RUnlock()

	preparedOrders := make([]*gen.PreparedOrder, len(s.orderIdToItems))
	for orderId, items := range s.orderIdToItems {
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

func (s *state) GetItemInfo(item *gen.Item) (ItemInfoModel, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	for orderId, items := range s.orderIdToItems {
		for i, itm := range items {
			itemsEqual, err := itemsEqual(item, itm)
			if err != nil {
				return ItemInfoModel{}, errors.New(err.Error())
			}
			if itemsEqual {
				return ItemInfoModel{
					OrderId: orderId,
					Index:   i,
					CubbyId: s.orderIdToCubbyId[orderId],
				}, nil
			}
		}
	}

	return ItemInfoModel{}, errors.New("there is no such item in any of the loaded orders")
}

func (s *state) RemoveItemFromOrder(orderId string, itemIndex int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.orderIdToItems[orderId][itemIndex] = s.orderIdToItems[orderId][len(s.orderIdToItems[orderId])-1]
	s.orderIdToItems[orderId] = s.orderIdToItems[orderId][:len(s.orderIdToItems[orderId])-1]
	return nil // think what could go wrong in this method
}

func (s *state) GetRemainingItemsCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	count := 0
	for _, items := range s.orderIdToItems {
		count += len(items)
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
