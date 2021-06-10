package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{}
}

type sortingService struct {
	items        []*gen.Item
	itemSelected *gen.Item
}

var randSource = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randSource)

func (s *sortingService) LoadItems(ctx context.Context, reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	s.items = append(s.items, reqPayload.Items...)

	copy(s.items, reqPayload.Items)

	log.Println("items loaded: ")

	s.printItems()

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) SelectItem(ctx context.Context, reqPayload *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	if s.itemSelected != nil {
		return nil, errors.New("item already selected in hand")
	}

	if len(s.items) == 0 {
		return nil, errors.New("no items to select")
	}

	itemsCount := len(s.items)

	var randomItemIndex int

	if itemsCount == 1 {
		randomItemIndex = 0
	} else {
		randomItemIndex = random.Intn(itemsCount - 1)
	}
	// TODO: use mutex
	s.itemSelected = s.items[randomItemIndex]

	log.Printf("item picked: %v\n", s.itemSelected.Code)

	s.items[randomItemIndex] = s.items[itemsCount-1]

	s.items = s.items[:itemsCount-1]

	return &gen.SelectItemResponse{Item: s.itemSelected}, nil
}

func (s *sortingService) MoveItem(ctx context.Context, reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	if s.itemSelected == nil {
		return nil, errors.New("no item in hand")
	}

	log.Printf("item %v moved to %v\n", s.itemSelected.Code, reqPayload.Cubby.Id)
	log.Println("items remaining:")
	s.printItems()

	s.itemSelected = nil

	return &gen.MoveItemResponse{}, nil
}

func (s *sortingService) printItems() {
	for _, v := range s.items {
		log.Printf("\t%v\n", v)
	}
}
