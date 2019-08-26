package main

import (
	"log"

	"github.com/go-chi/chi"
	"github.com/rahardiandj/go-transport/graphql/postgres"
)

func main() {

}

func initializeAPI() {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "postgres", "admin", "go_graphql_db"),
	)

	if err != nil {
		log.Fatalf("could not connect to postgres : %v", err)
	}

}
