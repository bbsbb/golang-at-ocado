package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	return &sortingService{
		cubbyToItems: make(map[string][]*gen.Item),
	}
}

type sortingService struct {
	mu           sync.Mutex
	bin          []*gen.Item
	itemSelected *gen.Item
	cubbyToItems map[string][]*gen.Item
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

func (s *sortingService) RemoveItemsByCode(ctx context.Context, reqPayload *gen.RemoveItemsRequest) (*gen.RemoveItemsResponse, error) {
	return s.removeItemsByCode(reqPayload)
}

func (s *sortingService) AuditState(ctx context.Context, in *gen.AuditStateRequest) (*gen.AuditStateResponse, error) {
	return s.auditState(in)
}

func (s *sortingService) loadItems(reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.bin = append(s.bin, reqPayload.Items...)

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) selectItem() (*gen.SelectItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.itemSelected != nil {
		return nil, errors.New("item already selected in hand")
	}

	if len(s.bin) == 0 {
		return nil, errors.New("no items to select")
	}

	itemsCount := len(s.bin)
	randomItemIndex := 0

	if itemsCount > 1 {
		randomItemIndex = random.Intn(itemsCount - 1)
	}

	fmt.Println("picking item...")
	//<-time.After(1 * time.Second)

	s.itemSelected = s.bin[randomItemIndex]
	s.bin[randomItemIndex] = s.bin[itemsCount-1]
	s.bin = s.bin[:itemsCount-1]

	return &gen.SelectItemResponse{Item: s.itemSelected}, nil
}

func (s *sortingService) moveItem(reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.itemSelected == nil {
		return nil, errors.New("no item in hand")
	}
	if !isValidCubbyID(reqPayload.Cubby.Id) {
		return nil, errors.New("invalid cubbyId - only 1 - 10 for now")
	}

	fmt.Println("moving item...")
	//<-time.After(1 * time.Second)

	s.cubbyToItems[reqPayload.Cubby.Id] = append(s.cubbyToItems[reqPayload.Cubby.Id], s.itemSelected)
	s.itemSelected = nil

	return &gen.MoveItemResponse{}, nil
}

func (s *sortingService) removeItemsByCode(reqPayload *gen.RemoveItemsRequest) (*gen.RemoveItemsResponse, error) {
	log.Printf("Removing [%d] items from the Bin", len(reqPayload.ItemCodes))
	removed := 0
	for _, code := range reqPayload.ItemCodes {
		for idx, item := range s.bin {
			if item.Code == code {
				s.bin = append(s.bin[:idx], s.bin[idx+1:]...)
				removed++
				break
			}
		}
	}

	log.Printf("Removed [%d] items while skipping [%d]", removed, len(reqPayload.ItemCodes)-removed)

	return &gen.RemoveItemsResponse{}, nil
}

func (s *sortingService) auditState(in *gen.AuditStateRequest) (*gen.AuditStateResponse, error) {
	cubbiesToItems := []*gen.CubbyToItems{}
	for cubby, items := range s.cubbyToItems {
		cubbiesToItems = append(cubbiesToItems, &gen.CubbyToItems{
			Cubby: &gen.Cubby{Id: cubby},
			Items: items,
		})
	}

	return &gen.AuditStateResponse{CubbiesToItems: cubbiesToItems}, nil
}

func isValidCubbyID(id string) bool {
	n, err := strconv.Atoi(id)
	return err == nil && n >= 1 && n <= 10
}
