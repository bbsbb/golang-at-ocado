package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
	"google.golang.org/grpc"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{
		orderIdItems:   make(map[string][]*gen.Item),
		orderIdCubbyId: make(map[string]string),
	}
}

type fulfillmentService struct {
	orderIdItems   map[string][]*gen.Item
	orderIdCubbyId map[string]string
}

func (s *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	preparedOrders := []*gen.PreparedOrder{}

	for _, order := range in.Orders {
		cubbyIdForOrder := s.determineCubbyIdForOrder(order)

		s.orderIdCubbyId[order.Id] = cubbyIdForOrder
		fmt.Printf("order: %v -> cubby: %v\n", order.Id, cubbyIdForOrder)

		s.orderIdItems[order.Id] = make([]*gen.Item, len(order.Items))
		copy(s.orderIdItems[order.Id], order.Items)

		preparedOrders = append(preparedOrders, &gen.PreparedOrder{
			Order: order,
			Cubby: &gen.Cubby{
				Id: cubbyIdForOrder,
			},
		})
	}

	s.commandSotingRobot()

	return &gen.CompleteResponse{
		Status: "ok",
		Orders: preparedOrders,
	}, nil
}

func (s *fulfillmentService) runSelectItem(client gen.SortingRobotClient) (*gen.SelectItemResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SelectItem(ctx, &gen.SelectItemRequest{})
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return res, nil
}

func (s *fulfillmentService) runMoveItem(client gen.SortingRobotClient, item *gen.Item) (*gen.MoveItemResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	orderId, itemIndexInOrder, cubbyId, err := s.getItemCubbyId(item)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	moveItemRequestPayload := &gen.MoveItemRequest{
		Cubby: &gen.Cubby{
			Id: cubbyId,
		},
	}
	res, err := client.MoveItem(ctx, moveItemRequestPayload)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	s.removeItemFromOrder(itemIndexInOrder, orderId)

	return res, nil
}

func (s *fulfillmentService) cubbyIsOccupied(cubby string) bool {
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

func (s *fulfillmentService) determineCubbyIdForOrder(order *gen.Order) string {
	cubbyId := ordertocubby.Map(order.Id, 1, 10)
	for times := 2; s.cubbyIsOccupied(cubbyId); times++ {
		cubbyId = ordertocubby.Map(order.Id, uint32(times), 10)
	}

	return cubbyId
}

func (s *fulfillmentService) getItemCubbyId(item *gen.Item) (string, int, string, error) {
	for orderId, items := range s.orderIdItems {
		for i, itm := range items {
			if item.Code == itm.Code {
				return orderId, i, s.orderIdCubbyId[orderId], nil
			}
		}
	}

	return "", -1, "", errors.New("there is no such item in any of the loaded orders")
}

func (s *fulfillmentService) getRemainingItemsCount() int {
	count := 0
	for _, items := range s.orderIdItems {
		count += len(items)
	}
	return count
}

func (s *fulfillmentService) removeItemFromOrder(itemIndex int, orderId string) {
	s.orderIdItems[orderId][itemIndex] = s.orderIdItems[orderId][len(s.orderIdItems[orderId])-1]
	s.orderIdItems[orderId] = s.orderIdItems[orderId][:len(s.orderIdItems[orderId])-1]
}

func (s *fulfillmentService) commandSotingRobot() {
	conn, client, err := getSortingRobotClient()
	defer conn.Close()

	if err != nil {
		// handle error
	}

	for s.getRemainingItemsCount() > 0 {
		res, err := s.runSelectItem(client)
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}
		_, err = s.runMoveItem(client, res.Item)
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}
	}
}

func getSortingRobotClient() (*grpc.ClientConn, gen.SortingRobotClient, error) {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, nil, errors.New(err.Error())
	}

	return conn, gen.NewSortingRobotClient(conn), nil
}
