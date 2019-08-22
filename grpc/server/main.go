package main

import (
	"fmt"
	"log"
	"net"

	api "github.com/rahardiandj/go-transport/grpc"

	"google.golang.org/grpc"
)

func main() {
	int port = 8889

	list, err := net.Listen("tcp", fmt.Sprintf(":%d", 8889))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := api.Server{}

	grpcServer := grpc.NewServer()

	api.RegisterCheckServer(grpcServer, &serv)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to start grpc server : %v", err)
	}

	log.Printf("")

}
