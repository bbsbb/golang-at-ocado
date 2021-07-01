package main

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/bbsbb/go-at-ocado/sort/gen"
	"github.com/preslavmihaylov/ordertocubby"
)

const cubbiesCnt = 10

func newFulfillmentService(client gen.SortingRobotClient) gen.FulfillmentServer {
	fs := &fulfillmentService{
		sortingRobot: client,
		orders:       make(map[string]*gen.FullfillmentStatus),
	}
	fs.ordersCh = scheduleWork(fs.processBatch)
	return fs
}

type fulfillmentService struct {
	sortingRobot      gen.SortingRobotClient
	ordersCh          chan []*gen.Order
	orders            map[string]*gen.FullfillmentStatus
	totallyConcurrent sync.Mutex
}

func (fs *fulfillmentService) processBatch(orders []*gen.Order) {

	for k := range fs.orders {
		log.Println(k)
	}

	// ordersToCubbies := mapOrdersToCubbies(orders)

	// for _, order := range orders {
	// 	for _, item := range order.Items {
	// 		_ = item

	// 		resp, err := fs.sortingRobot.PickItem(context.Background(), &gen.Empty{})
	// 		if err != nil {
	// 			// What do?
	// 		}

	// 		cubbyID := getCubbyForItem(resp.Item, orders, ordersToCubbies)
	// 		_, err = fs.sortingRobot.PlaceInCubby(context.Background(), &gen.PlaceInCubbyRequest{
	// 			Cubby: &gen.Cubby{Id: cubbyID},
	// 		})
	// 		if err != nil {
	// 			// What do?
	// 		}
	// 	}
	// }
}

func (fs *fulfillmentService) LoadOrders(ctx context.Context, in *gen.LoadOrdersRequest) (*gen.CompleteResponse, error) {
	go func() {
		fs.totallyConcurrent.Lock()
		for _, o := range in.Orders {
			fs.orders[o.Id] = &gen.FullfillmentStatus{
				Order: o,
				Cubby: &gen.Cubby{},
				State: gen.OrderState_PENDING,
			}
		}
		fs.totallyConcurrent.Unlock()
		fs.ordersCh <- in.Orders
	}()

	return &gen.CompleteResponse{}, nil
}

func mapOrdersToCubbies(orders []*gen.Order) map[string]string {
	ordersToCubbies := map[string]string{}
	usedCubbies := map[string]bool{}

	for _, order := range orders {
		cubbyID := mapOrderToCubby(usedCubbies, order.Id, cubbiesCnt)
		ordersToCubbies[order.Id] = cubbyID
		usedCubbies[cubbyID] = true
	}

	for orderID, cubbyID := range ordersToCubbies {
		fmt.Printf("order %s -> cubby %s\n", orderID, cubbyID)
	}

	return ordersToCubbies
}

func mapOrderToCubby(usedCubbies map[string]bool, id string, cubbiesCnt int) string {
	times := 1
	for {
		cubbyID := ordertocubby.Map(id, uint32(times), uint32(cubbiesCnt))
		if !usedCubbies[cubbyID] {
			return cubbyID
		}

		times++
	}
}

func getOrderForItem(item *gen.Item, orders []*gen.Order) *gen.Order {
	var match *gen.Order
	return match
}

func getCubbyForItem(item *gen.Item, orders []*gen.Order, cubbyMapping map[string]string) string {
	// We get an item
	// We solve it to order - handle item in two possible orders
	// We lookup the cubbyId for the orders

	log.Println(mapOrdersToCubbies)
	log.Fatal("DONE")
	return "1"
}

func scheduleWork(work func([]*gen.Order)) chan []*gen.Order {
	ordersCh := make(chan []*gen.Order)
	go func() {
		log.Printf("Initializing orders worker...")
		for {
			orders := <-ordersCh
			work(orders)
		}
	}()
	return ordersCh
}

func (f *fulfillmentService) GetOrderStatusById(ctx context.Context, in *gen.OrderIdRequest) (*gen.OrdersStatusResponse, error) {
	return nil, nil
}

func (f *fulfillmentService) GetAllOrdersStatus(context.Context, *gen.Empty) (*gen.OrdersStatusResponse, error) {
	return nil, nil
}

func (f *fulfillmentService) MarkFullfilled(context.Context, *gen.OrderIdRequest) (*gen.Empty, error) {
	return nil, nil
}
