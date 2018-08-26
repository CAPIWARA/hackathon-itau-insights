package lgq

import (
	"github.com/graphql-go/graphql"
)

var fieldGetUser = &graphql.Field{
	Type:        typeUser,
	Description: "get user",
	Resolve:     GetUser,
}

var fieldGetProject = &graphql.Field{
	Type:        typeProject,
	Description: "get project",
	Args: graphql.FieldConfigArgument{
		"projectId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: GetProject,
}

var fieldCreateProject = &graphql.Field{
	Type:        typeProject,
	Description: "create project",
	Resolve:     CreateProject,
}

var fieldPayment = &graphql.Field{
	Type:        graphql.NewList(typePayment),
	Description: "get payments by project",
	Resolve:     GetPayments,
}
