package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newSortingService() gen.SortingServer {
	rand.Seed(time.Now().UnixNano())
	return &sortingService{}
}

type sortingService struct {
	Bin        []*gen.Item
	pickedItem *gen.Item
}

func (s *sortingService) LoadItems(ctx context.Context, req *gen.LoadItemsRequest) (*gen.Empty, error) {
	s.Bin = append(s.Bin, req.Items...)
	log.Printf("Added [%d] items to the bin, for total storage of [%d]", len(req.Items), len(s.Bin))

	return &gen.Empty{}, nil
}

func (s *sortingService) PickItem(context.Context, *gen.Empty) (*gen.PickItemResponse, error) {
	if len(s.Bin) < 1 {
		log.Println("no items in the bin, get out.")
		return nil, errors.New("no items in the bin, get out")
	} else if s.pickedItem != nil {
		log.Println("an item is already picked")
		return nil, errors.New("an item is already picked")
	}

	itemPos := rand.Intn(len(s.Bin))
	item := s.Bin[itemPos]
	s.pickedItem = item
	s.Bin = append(s.Bin[:itemPos], s.Bin[itemPos+1:]...)

	log.Printf("Picked item at position [%d] from the bin, [%d] items left", itemPos, len(s.Bin))
	return &gen.PickItemResponse{
		Item: item,
	}, nil
}

func (s *sortingService) PlaceInCubby(context.Context, *gen.PlaceInCubbyRequest) (*gen.Empty, error) {
	if s.pickedItem == nil {
		log.Println("no item is currently picked")
		return nil, errors.New("no item is currently picked")
	}

	s.pickedItem = nil
	return &gen.Empty{}, nil
}

func (s *sortingService) AuditState(context.Context, *gen.Empty) (*gen.AuditStateResponse, error) {
	return nil, errors.New("not implemented")
}
