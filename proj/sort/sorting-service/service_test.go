package main

import (
	"testing"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"github.com/stretchr/testify/assert"
)

func TestLoadItems(t *testing.T) {
	testCases := []struct {
		name                string
		service             *sortingService
		loadItemsReqPayload *gen.LoadItemsRequest
	}{
		{
			name:                "test load items",
			service:             getEmptyService(),
			loadItemsReqPayload: getLoadItemsPayload(),
		},
		{
			name:                "test load items on non empty main cubby",
			service:             getLoadedService(),
			loadItemsReqPayload: getLoadItemsPayload(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			initialItemsCount := len(tc.service.bin)
			addedItemsCount := len(tc.loadItemsReqPayload.Items)

			_, err := tc.service.loadItems(tc.loadItemsReqPayload)
			assert.Nil(t, err)
			assert.Equal(t, initialItemsCount+addedItemsCount, len(tc.service.bin))
		})
	}
}

func TestSelectItem(t *testing.T) {
	service := getLoadedService()
	itemsCountBeforeSelecting := len(service.bin)

	res, err := service.selectItem()
	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, len(service.bin), itemsCountBeforeSelecting-1)
}

func TestSelectItemErrors(t *testing.T) {
	testCases := []struct {
		name    string
		service *sortingService
	}{
		{
			name:    "test select item on empty main cubby",
			service: getEmptyService(),
		},
		{
			name:    "test select item already selected item in hand",
			service: getLoadedServiceAndSelect(),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := tc.service.selectItem()
			assert.NotNil(t, err)
		})
	}
}

func TestMoveItem(t *testing.T) {
	loadedService := getLoadedServiceAndSelect()

	_, err := loadedService.moveItem(getMoveItemsPayload())
	assert.Nil(t, err)
	assert.Nil(t, loadedService.itemSelected)
}

func TestMoveItemEmptyHandError(t *testing.T) {
	service := getEmptyService()

	_, err := service.moveItem(getMoveItemsPayload())
	assert.NotNil(t, err)
}
