package gql

import (
	"github.com/graphql-go/graphql"
	"github.com/rahardiandj/go-transport/graphql/postgres"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {

	resolver := Resolver{}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolver.UserResolver,
					},
				},
			},
		),
	}
	return &root
}
