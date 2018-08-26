package lgq

import (
	"github.com/graphql-go/graphql"
)

var typeUser = graphql.NewObject(graphql.ObjectConfig{
	Name: "typeUser",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Name: "name",
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Name: "email",
			Type: graphql.String,
		},
	},
})

var typeProject = graphql.NewObject(graphql.ObjectConfig{
	Name: "typeProject",
	Fields: graphql.Fields{
		"createdBy": &graphql.Field{
			Name: "createdBy",
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Name: "_id",
			Type: graphql.String,
		},
	},
})

var typePayment = graphql.NewObject(graphql.ObjectConfig{
	Name: "typePayment",
	Fields: graphql.Fields{
		"transactionId": &graphql.Field{
			Name: "transactionId",
			Type: graphql.String,
		},
		"projectId": &graphql.Field{
			Name: "projectId",
			Type: graphql.String,
		},
		"userId": &graphql.Field{
			Name: "userId",
			Type: graphql.String,
		},
		"id": &graphql.Field{
			Name: "_id",
			Type: graphql.String,
		},
	}})
