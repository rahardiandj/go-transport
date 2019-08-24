package main

import (
	"context"
	"log"

	api "github.com/rahardiandj/go-transport/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var conn *grpc.ClientConn

	creds, err := credentials.NewClientTLSFromFile("../cert/server.crt", "")
	if err != nil {
		log.Fatalf("could not load tls cert: %s", err)
	}

	conn, err = grpc.Dial(":8889", grpc.WithTransportCredentials(creds))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := api.NewCheckClient(conn)

	response, err := c.Ping(context.Background(), &api.Void{})

	if err != nil {
		log.Printf("Error when do health chec: %v", err)
	}

	log.Printf("Health check status : %s", response.Check)
}
