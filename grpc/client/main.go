package main

import (
	"context"
	"log"

	api "github.com/rahardiandj/go-transport/grpc"

	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":8889", grpc.WithInsecure())

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
