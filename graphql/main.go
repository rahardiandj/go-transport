package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/rahardiandj/go-transport/graphql/server"

	"github.com/graphql-go/graphql"

	"github.com/go-chi/chi"
	"github.com/rahardiandj/go-transport/graphql/gql"
	"github.com/rahardiandj/go-transport/graphql/postgres"
)

func main() {

	router, db := initializeAPI()
	defer db.Close()

	log.Fatal(http.ListenAndServe(":8889", router))
}

func initializeAPI() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(
		postgres.ConnString("localhost", 5432, "postgres", "postgres", "go_graphql_db"),
	)

	if err != nil {
		log.Fatalf("could not connect to postgres : %v", err)
	}

	rootQuery := gql.NewRoot(db)

	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: rootQuery.Query},
	)

	if err != nil {
		fmt.Printf("error creating schema: %v\n", err)
	}

	s := server.Server{
		GqlSchema: &sc,
	}

	//add middleware to the router
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.StripSlashes,
		middleware.Recoverer, // recover from panic witout crashing server
	)

	router.Post("/graphql", s.GraphQL())

	return router, db

}
