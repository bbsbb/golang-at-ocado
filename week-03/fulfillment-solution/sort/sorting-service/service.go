package main

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	rand.Seed(time.Now().UnixNano())
	return &sortingService{
		cubbyToItems: map[string][]*gen.Item{},
	}
}

type sortingService struct {
	Bin          []*gen.Item
	pickedItem   *gen.Item
	cubbyToItems map[string][]*gen.Item
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

func (s *sortingService) PlaceInCubby(ctx context.Context, req *gen.PlaceInCubbyRequest) (*gen.Empty, error) {
	if s.pickedItem == nil {
		log.Println("no item is currently picked")
		return nil, errors.New("no item is currently picked")
	} else if !isValidCubbyID(req.Cubby.Id) {
		log.Printf("received invalid cubby id: %s", req.Cubby.Id)
		return nil, errors.New("invalid cubby ID. Should be in range [1..10]")
	}

	log.Printf("Placed %s in cubby %s", s.pickedItem.Code, req.Cubby.Id)
	s.cubbyToItems[req.Cubby.Id] = append(s.cubbyToItems[req.Cubby.Id], s.pickedItem)
	s.pickedItem = nil

	return &gen.Empty{}, nil
}

func (s *sortingService) AuditState(context.Context, *gen.Empty) (*gen.AuditStateResponse, error) {
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
