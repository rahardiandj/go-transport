package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc/credentials"

	api "github.com/rahardiandj/go-transport/grpc"

	"google.golang.org/grpc"
)

func main() {
	port := 8889

	list, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	serv := api.Server{}

	creds, err := credentials.NewServerTLSFromFile("../cert/server.crt", "../cert/server.key")

	if err != nil {
		log.Fatalf("could not load TLS keys: %v", err)
	}

	opts := []grpc.ServerOption{grpc.Creds(creds)}

	grpcServer := grpc.NewServer(opts...)

	api.RegisterCheckServer(grpcServer, &serv)

	log.Printf("Staring grpc server with port: %d ....", port)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to start grpc server : %v", err)
	}

}
