package main

import "github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"

func getEmptyService() *sortingService {
	return &sortingService{}
}

func getLoadedService() *sortingService {
	s := getEmptyService()
	s.loadItems(getLoadItemsPayload())

	return s
}

func getLoadedServiceAndSelect() *sortingService {
	s := getLoadedService()
	s.selectItem()

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
