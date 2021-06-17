package main

import (
	"context"
	"errors"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{}
}

type fulfillmentService struct {
}

func (s *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	return nil, errors.New("not implemented")
}
