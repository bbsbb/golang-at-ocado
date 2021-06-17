package main

import (
	"context"
	"errors"

	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newFulfillmentService() gen.FulfillmentServer {
	return &fulfillmentService{}
}

type fulfillmentService struct{}

func (fs *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	return nil, errors.New("not implemented")
}
