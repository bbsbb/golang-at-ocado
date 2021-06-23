package main

import (
	"context"
	"errors"
	"fmt"

	"github.com/bbsbb/go-at-ocado/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
)

const cubbiesCnt = 10

func newFulfillmentService(client gen.SortingRobotClient) gen.FulfillmentServer {
	return &fulfillmentService{
		sortingRobot: client,
	}
}

type fulfillmentService struct {
	sortingRobot gen.SortingRobotClient
}

func (fs *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	// map orders to cubbies
	ordersToCubbies := mapOrdersToCubbies(in.Orders)
	_ = ordersToCubbies

	for _, order := range in.Orders {
		for _, item := range order.Items {
			_ = item

			resp, err := fs.sortingRobot.PickItem(ctx, &gen.Empty{})
			if err != nil {
				return nil, &FulfillmentFailedError{}
			}

			cubbyID := getCubbyForItem(resp.Item)
			_, err = fs.sortingRobot.PlaceInCubby(ctx, &gen.PlaceInCubbyRequest{
				Cubby: &gen.Cubby{Id: cubbyID},
			})
			if err != nil {
				return nil, fmt.Errorf("place in cubby failed: %v", err)
			}
		}
	}

	return nil, errors.New("not implemented")
}

func mapOrdersToCubbies(orders []*gen.Order) map[string]string {
	ordersToCubbies := map[string]string{}
	usedCubbies := map[string]bool{}

	for _, order := range orders {
		cubbyID := mapOrderToCubby(usedCubbies, order.Id, cubbiesCnt)
		ordersToCubbies[order.Id] = cubbyID
		usedCubbies[cubbyID] = true
	}

	for orderID, cubbyID := range ordersToCubbies {
		fmt.Printf("order %s -> cubby %s\n", orderID, cubbyID)
	}

	return ordersToCubbies
}

func mapOrderToCubby(usedCubbies map[string]bool, id string, cubbiesCnt int) string {
	times := 1
	for {
		cubbyID := ordertocubby.Map(id, uint32(times), uint32(cubbiesCnt))
		if !usedCubbies[cubbyID] {
			return cubbyID
		}

		times++
	}
}

func getCubbyForItem(item *gen.Item) string {
	return "1"
}
