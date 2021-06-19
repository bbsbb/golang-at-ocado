package main

import (
	"context"
	"errors"
	"math/rand"
	"sync"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{}
}

type sortingService struct {
	mu           sync.Mutex
	items        []*gen.Item
	itemSelected *gen.Item
}

var randSource = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randSource)

func (s *sortingService) LoadItems(ctx context.Context, reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	return s.loadItems(reqPayload)
}

func (s *sortingService) SelectItem(ctx context.Context, reqPayload *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	return s.selectItem()
}

func (s *sortingService) MoveItem(ctx context.Context, reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	return s.moveItem(reqPayload)
}

func (s *sortingService) loadItems(reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.items = append(s.items, reqPayload.Items...)

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) selectItem() (*gen.SelectItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.itemSelected != nil {
		return nil, errors.New("item already selected in hand")
	}

	if len(s.items) == 0 {
		return nil, errors.New("no items to select")
	}

	itemsCount := len(s.items)
	randomItemIndex := 0

	if itemsCount > 1 {
		randomItemIndex = random.Intn(itemsCount - 1)
	}

	s.itemSelected = s.items[randomItemIndex]
	s.items[randomItemIndex] = s.items[itemsCount-1]
	s.items = s.items[:itemsCount-1]

	return &gen.SelectItemResponse{Item: s.itemSelected}, nil
}

func (s *sortingService) moveItem(reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.itemSelected == nil {
		return nil, errors.New("no item in hand")
	}
	s.itemSelected = nil

	return &gen.MoveItemResponse{}, nil
}
