package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{expectedNextRequest: load}
}

type sortingService struct {
	items               []*gen.Item
	expectedNextRequest string
	itemSelected        *gen.Item
}

// Expected next request types
const (
	load  = "load"
	selct = "select"
	move  = "move"
)

var randSource = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randSource)

func (s *sortingService) LoadItems(ctx context.Context, reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	if s.expectedNextRequest != load {
		return nil, errors.New(fmt.Sprintf("expected %v", s.expectedNextRequest))
	}

	s.items = make([]*gen.Item, len(reqPayload.Items))

	copy(s.items, reqPayload.Items)

	log.Println("loaded items:")

	s.printItems()

	s.expectedNextRequest = selct

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) SelectItem(ctx context.Context, reqPayload *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	if s.expectedNextRequest != selct {
		return nil, s.getError()
	}

	itemsCount := len(s.items)

	var randomItemIndex int

	if itemsCount == 1 {
		randomItemIndex = 0
	} else {
		randomItemIndex = random.Intn(itemsCount - 1)
	}

	s.itemSelected = s.items[randomItemIndex]

	log.Printf("item picked: %v\n", s.itemSelected.Code)

	s.items[randomItemIndex] = s.items[itemsCount-1]

	s.items = s.items[:itemsCount-1]

	s.expectedNextRequest = move

	return &gen.SelectItemResponse{Item: s.itemSelected}, nil
}

func (s *sortingService) MoveItem(ctx context.Context, reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	if s.expectedNextRequest != move {
		return nil, s.getError()
	}

	log.Printf("item %v moved to %v\n", s.itemSelected.Code, reqPayload.Cubby.Id)
	log.Println("items remaining:")
	s.printItems()

	s.itemSelected = nil

	if len(s.items) == 0 {
		s.expectedNextRequest = load
	} else {
		s.expectedNextRequest = selct
	}

	return &gen.MoveItemResponse{}, nil
}

func (s *sortingService) getError() error {
	return errors.New(fmt.Sprintf("expected %v", s.expectedNextRequest))
}

func (s *sortingService) printItems() {
	for _, v := range s.items {
		log.Printf("\t%v\n", v)
	}
}
