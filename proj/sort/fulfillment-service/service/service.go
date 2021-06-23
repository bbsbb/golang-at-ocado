package service

import (
	"context"
	"fmt"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/fulfillment-service/state"
	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

type fulfillmentService struct {
	state              state.State
	sortingRobotClient gen.SortingRobotClient
}

func New(sortingRobotClient gen.SortingRobotClient) gen.FulfillmentServer {
	return &fulfillmentService{
		state:              state.New(),
		sortingRobotClient: sortingRobotClient,
	}
}

func (s *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	err := s.state.PersistOrders(in.Orders)
	if err != nil {
		return nil, err
	}
	preparedOrders := s.state.GetPreparedOrders()
	s.sortItemsInCubbies(ctx)

	return &gen.CompleteResponse{Status: "ok", Orders: preparedOrders}, nil
}

func (s *fulfillmentService) sortItemsInCubbies(ctx context.Context) {
	for s.state.GetRemainingItemsCount() > 0 {
		res, err := s.sortingRobotClient.SelectItem(ctx, &gen.SelectItemRequest{})
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}
		itemInfo, err := s.state.GetItemInfo(res.Item)
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}
		_, err = s.sortingRobotClient.MoveItem(ctx, &gen.MoveItemRequest{Cubby: &gen.Cubby{Id: itemInfo.CubbyId}})
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}
		s.state.RemoveItemFromOrder(itemInfo.OrderId, itemInfo.Index)
		// Eventually think out to mark item as 'added to cubby' and clear order if all items are already in a cubby
	}
}
