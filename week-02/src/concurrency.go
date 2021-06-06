package main

import (
	"fmt"
	"log"
	"time"
)

// Resources:
// https://blog.golang.org/pipelines
// https://blog.golang.org/context
// https://www.youtube.com/watch?v=oV9rvDllKEg
// x/errgrp - https://www.fullstory.com/blog/why-errgroup-withcontext-in-golang-server-handlers
// Scheduler: https://github.com/golang/go/blob/master/src/runtime/proc.go

func orderline(payments []int, outputCh chan<- int, delay uint32) {
	totalDue := 0
	for _, due := range payments {
		totalDue += due
	}

	time.Sleep(time.Duration(delay) * time.Second)
	outputCh <- totalDue
}

func main() {
	start := time.Now()

	payments := []int{1, 3, 5, 7, 9, 13}
	resultCh := make(chan int)

	go orderline(payments[len(payments)-1:], resultCh, 5)
	go orderline(payments[:len(payments)-1], resultCh, 3)

	batchOne, batchTwo := <-resultCh, <-resultCh

	fmt.Printf("Total is: %d\n in %v", batchOne+batchTwo, time.Since(start))
}
