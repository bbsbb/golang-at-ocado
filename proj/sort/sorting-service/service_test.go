package main

import (
	"context"
	"testing"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func TestLoadItems(t *testing.T) {
	service := &sortingService{}

	initialLength := len(service.items)

	loadItemsReqPayload := getLoadItemsPayload()

	addedItemsCount := len(loadItemsReqPayload.Items)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	service.LoadItems(ctx, loadItemsReqPayload)

	if initialLength+addedItemsCount != len(service.items) {
		t.Error("length of payload != length of loaded items")
		t.FailNow()
	}
}

func TestSelectItem(t *testing.T) {
	loadedService := getLoadedService()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	itemsCountBeforeSelecting := len(loadedService.items)

	res, _ := loadedService.SelectItem(ctx, &gen.SelectItemRequest{})

	if res == nil {
		t.Error("no item picked")
		t.FailNow()
	}

	if itemsCountBeforeSelecting != len(loadedService.items)+1 {
		t.Error("number of item not reduced by one")
		t.FailNow()
	}
}

func TestSelectItemEmptyArrayError(t *testing.T) {
	loadedService := &sortingService{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	_, err := loadedService.SelectItem(ctx, &gen.SelectItemRequest{})

	if err == nil {
		t.Error("error expected")
		t.FailNow()
	}
}

func TestSelectItemAlreadySelectedItemError(t *testing.T) {
	loadedService := getLoadedService()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	loadedService.SelectItem(ctx, &gen.SelectItemRequest{})

	_, err := loadedService.SelectItem(ctx, &gen.SelectItemRequest{})

	if err == nil {
		t.Error("error expected")
		t.FailNow()
	}
}

func getLoadedService() *sortingService {
	s := &sortingService{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	s.LoadItems(ctx, getLoadItemsPayload())

	return s
}

func getLoadItemsPayload() *gen.LoadItemsRequest {
	return &gen.LoadItemsRequest{Items: []*gen.Item{
		&gen.Item{Code: "1234", Label: "label1"},
		&gen.Item{Code: "2345", Label: "label2"},
		&gen.Item{Code: "3456", Label: "label3"},
	}}
}
