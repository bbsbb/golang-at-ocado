package main

import (
	"context"
	"errors"

	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{}
}

type sortingService struct{}

func (s *sortingService) LoadItems(context.Context, *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	// TODO: Implement this
	return nil, errors.New("not implemented")
}

func (s *sortingService) MoveItem(context.Context, *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	// TODO: Implement this
	return nil, errors.New("not implemented")
}

func (s *sortingService) SelectItem(context.Context, *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	// TODO: Implement this
	return nil, errors.New("not implemented")
}
