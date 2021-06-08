package main

import (
	"github.com/bbsbb/go-at-ocado/sort/gen"
)

func newSortingService() gen.SortingServer {
	// TODO: Implement the gRPC interface
	return &sortingService{}
}

type sortingService struct{}
