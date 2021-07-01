package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/fulfillment-service/state"
	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

// TODO: figure out how to extract order processor in another object
type FulfillmentService interface {
	LoadOrders(context.Context, *gen.LoadOrdersRequest) (*gen.CompleteResponse, error)
	GetOrderStatusById(context.Context, *gen.OrderIdRequest) (*gen.OrdersStatusResponse, error)
	GetAllOrdersStatus(context.Context, *gen.OrdersStatusRequest) (*gen.OrdersStatusResponse, error)
	MarkFullfilled(context.Context, *gen.OrderIdRequest) (*gen.MarkFullfilledResponse, error)
	StartExpectingWork()
}

// TODO: if it's a lot FulfilmentServiceParameters
func New(sortingRobotClient gen.SortingRobotClient) FulfillmentService {
	return &fulfillmentService{
		state:              state.New(), // TODO: pass it 
		sortingRobotClient: sortingRobotClient,
		qChan:              make(chan *gen.LoadOrdersRequest), // TODO: pass it 
	}
}

type fulfillmentService struct {
	state              state.State
	sortingRobotClient gen.SortingRobotClient
	qChan              chan *gen.LoadOrdersRequest // TODO: change name
}

func (s *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	currentlyProcessing := s.state.IsCurrentlyProcessing()
	// TODO: inline with anonimous
	go s.passWorkRequest(in)

	if currentlyProcessing {
		return &gen.CompleteResponse{Status: "will start processing shortly", Orders: []*gen.PreparedOrder{}}, nil
	}

	return &gen.CompleteResponse{Status: "starting immediately", Orders: []*gen.PreparedOrder{}}, nil
}

func (s *fulfillmentService) GetOrderStatusById(ctx context.Context, in *gen.OrderIdRequest) (*gen.OrdersStatusResponse, error) {
	orderInfo, err := s.state.GetOrderInfo(in.OrderId)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &gen.OrdersStatusResponse{
		Status: []*gen.FullfillmentStatus{
			{
				State: orderInfo.State,
				Order: orderInfo.Order,
				Cubby: orderInfo.Cubby,
			},
		},
	}, nil
}

func (s *fulfillmentService) GetAllOrdersStatus(ctx context.Context, in *gen.OrdersStatusRequest) (*gen.OrdersStatusResponse, error) {
	
	// TODO: make a method get order state(do this at the end cuz the code is nice even as it is)
	orderIds, err := s.state.GetAllOrderIds()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	status := make([]*gen.FullfillmentStatus, len(orderIds))
	for _, orderId := range orderIds {
		orderInfo, err := s.state.GetOrderInfo(orderId)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		status = append(status, &gen.FullfillmentStatus{
			State: orderInfo.State,
			Order: orderInfo.Order,
			Cubby: orderInfo.Cubby,
		})
	}

	return &gen.OrdersStatusResponse{Status: status}, nil
}

func (s *fulfillmentService) MarkFullfilled(ctx context.Context, in *gen.OrderIdRequest) (*gen.MarkFullfilledResponse, error) {
	return nil, nil
}

// TODO: consider this func as a candidate for the new processing order object that we talked at the start
func (s *fulfillmentService) StartExpectingWork() {
	for {
		in := <-s.qChan // TODO: check how to handle channel read errors
		s.state.SetCurrentlyProccessingTrue()

		// TODO: concurency issue -> someone may changr the curently procesing flag in between the two methods
		if err := s.state.PersistOrders(in.Orders); err != nil {
			fmt.Println(err.Error())
		}

		s.sortItemsInCubbies(context.Background())
		s.state.SetCurrentlyProccessingFalse()
	}
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

		err = s.state.MarkItemInCubby(itemInfo.OrderId, itemInfo.Index)
		if err != nil {
			fmt.Println(err.Error()) // handle error
		}

		fmt.Printf("moved item %v to cubby %v\n", res.Item.Label, itemInfo.CubbyId)
	}
}

// TODO: pass it as an anonomous func
func (s *fulfillmentService) passWorkRequest(in *gen.LoadOrdersRequest) {
	s.qChan <- in
}
