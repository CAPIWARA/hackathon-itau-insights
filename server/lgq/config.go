package lgq

import (
	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    query,
	Mutation: mutation,
})

var query = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "query",
		Description: "query",
		Fields:      graphql.Fields{},
	},
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "mutation",
		Description: "mutation",
		Fields:      graphql.Fields{}})
