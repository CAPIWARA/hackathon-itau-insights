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
		Name:        "rootQuery",
		Description: "rootQuery",
		Fields: graphql.Fields{
			"user":     fieldGetUser,
			"project":  fieldGetProject,
			"payments": fieldPayment,
		},
	},
)

var mutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "rootMutation",
		Description: "rootMutation",
		Fields: graphql.Fields{
			"createProject": fieldCreateProject,
			"checkout":      fieldCheckout,
		},
	},
)
