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
}
