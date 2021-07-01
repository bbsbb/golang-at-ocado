package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/fulfillment-service/service"
	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// TODO: do in a one constant statement
const serverPort = ":10001"
const sortingRobotServerUrl = "localhost:10000"

func main() {
	conn, sortingRobotClient := getSortingRobotClient()
	defer conn.Close()

	grpcServer, lis := newFulfillmentServer(sortingRobotClient)

	fmt.Printf("gRPC server started. Listening on %s\n", serverPort)
	grpcServer.Serve(lis)
}

func newFulfillmentServer(sortingRobotClient gen.SortingRobotClient) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	fulfillmentService := service.New(sortingRobotClient)
	go fulfillmentService.StartExpectingWork()
	gen.RegisterFulfillmentServer(grpcServer, fulfillmentService)
	reflection.Register(grpcServer)

	return grpcServer, lis
}

func getSortingRobotClient() (*grpc.ClientConn, gen.SortingRobotClient) {
	conn, err := grpc.Dial(sortingRobotServerUrl, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	return conn, gen.NewSortingRobotClient(conn)
}
