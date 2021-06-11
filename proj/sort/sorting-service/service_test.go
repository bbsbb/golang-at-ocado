package main

import (
	"testing"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func TestLoadItems(t *testing.T) {
	service := &sortingService{}

	initialItemsCount := len(service.items)

	loadItemsReqPayload := getLoadItemsPayload()

	addedItemsCount := len(loadItemsReqPayload.Items)

	_, err := service.loadItems(loadItemsReqPayload)

	if err != nil {
		t.Errorf("error lading items: %v", err)
		t.FailNow()
	}

	if initialItemsCount+addedItemsCount != len(service.items) {
		t.Error("length of payload != length of loaded items")
		t.FailNow()
	}
}

func TestLoadItemsOnNonEmptyItemSlice(t *testing.T) {
	loadedService := getLoadedService()

	initialItemsCount := len(loadedService.items)

	loadItemsReqPayload := getLoadItemsPayload()

	addedItemsCount := len(loadItemsReqPayload.Items)

	_, err := loadedService.loadItems(loadItemsReqPayload)

	if err != nil {
		t.Errorf("error lading items: %v", err)
		t.FailNow()
	}

	if initialItemsCount+addedItemsCount != len(loadedService.items) {
		t.Error("len(items) + len(itemsToAdd) != len(itemsAfterLoading)")
		t.FailNow()
	}
}

func TestSelectItem(t *testing.T) {
	loadedService := getLoadedService()

	itemsCountBeforeSelecting := len(loadedService.items)

	res, err := loadedService.selectItem()

	if err != nil {
		t.Errorf("error selecting item: %v", err)
		t.FailNow()
	}

	if res == nil {
		t.Error("no item picked")
		t.FailNow()
	}

	if itemsCountBeforeSelecting != len(loadedService.items)+1 {
		t.Error("count of items not reduced by one")
		t.FailNow()
	}
}

func TestSelectItemEmptyItemSliceError(t *testing.T) {
	loadedService := &sortingService{}

	_, err := loadedService.selectItem()

	if err == nil {
		t.Error("error expected")
		t.FailNow()
	}
}

func TestSelectItemAlreadySelectedItemError(t *testing.T) {
	loadedService := getLoadedService()

	loadedService.selectItem()

	_, err := loadedService.selectItem()

	if err == nil {
		t.Error("error expected")
		t.FailNow()
	}
}

func TestMoveItem(t *testing.T) {
	loadedService := getLoadedService()

	loadedService.selectItem()

	_, err := loadedService.moveItem(getMoveItemsPayload())

	if err != nil {
		t.Errorf("error moving item: %v", err)
		t.FailNow()
	}

	if loadedService.itemSelected != nil {
		t.Error("hand not empty")
		t.FailNow()
	}
}

func TestMoveItemEmptyHandError(t *testing.T) {
	service := &sortingService{}

	_, err := service.moveItem(getMoveItemsPayload())

	if err == nil {
		t.Error("error expected")
		t.FailNow()
	}
}

func getLoadedService() *sortingService {
	s := &sortingService{}

	s.loadItems(getLoadItemsPayload())

	return s
}

func getLoadItemsPayload() *gen.LoadItemsRequest {
	return &gen.LoadItemsRequest{Items: []*gen.Item{
		{Code: "1234", Label: "label1"},
		{Code: "2345", Label: "label2"},
		{Code: "3456", Label: "label3"},
	}}
}

func getMoveItemsPayload() *gen.MoveItemRequest {
	return &gen.MoveItemRequest{
		Cubby: &gen.Cubby{Id: "1"},
	}
}
