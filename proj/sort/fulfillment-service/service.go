package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/fulfillment-service/state"
	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"google.golang.org/grpc"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{
		state: state.NewState(),
	}
}

type fulfillmentService struct {
	state state.State
}

func (s *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	err := s.state.PersistOrders(in)
	if err != nil {
		return nil, err
	}
	response, _ := s.state.GetCompleteResponse() // error ignored

	s.commandSotingRobot()

	return response, nil
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

	itemInfo, err := s.state.GetItemInfo(item)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	moveItemRequestPayload := &gen.MoveItemRequest{
		Cubby: &gen.Cubby{
			Id: itemInfo.CubbyId,
		},
	}
	res, err := client.MoveItem(ctx, moveItemRequestPayload)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	s.state.RemoveItemFromOrder(itemInfo.OrderId, itemInfo.Index)

	return res, nil
}

func (s *fulfillmentService) commandSotingRobot() {
	conn, client, err := getSortingRobotClient()
	defer conn.Close()

	if err != nil {
		// handle error
	}

	for s.state.GetRemainingItemsCount() > 0 {
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
