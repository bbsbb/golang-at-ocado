package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bbsbb/go-at-ocado/sort/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverPort = "localhost:10001"
const sortingRobotPort = "localhost:10000"

func main() {
	client, conn := newSortingRobotClient()
	defer conn.Close()

	grpcServer, lis := newFulfillmentServer(client)

	fmt.Printf("gRPC server started. Listening on %s\n", serverPort)
	grpcServer.Serve(lis)
}

func newSortingRobotClient() (gen.SortingRobotClient, *grpc.ClientConn) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(sortingRobotPort, opts...)
	if err != nil {
		panic("couldn't start server: " + err.Error())
	}

	return gen.NewSortingRobotClient(conn), conn
}

func newFulfillmentServer(client gen.SortingRobotClient) (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterFulfillmentServer(grpcServer, newFulfillmentService(client))
	reflection.Register(grpcServer)

	return grpcServer, lis
}
