package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bbsbb/go-at-ocado/sort/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serverPort = "localhost:10000"

func main() {
	grpcServer, lis := newSortingServer()

	fmt.Printf("gRPC server started. Listening on %s\n", serverPort)
	grpcServer.Serve(lis)
}

func newSortingServer() (*grpc.Server, net.Listener) {
	lis, err := net.Listen("tcp", serverPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	gen.RegisterSortingRobotServer(grpcServer, newSortingService())
	reflection.Register(grpcServer)

	return grpcServer, lis
}
