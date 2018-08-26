package lgq

import (
	"hackathon-itau-insights/server/payment"
	"hackathon-itau-insights/server/project"
	"hackathon-itau-insights/server/user"

	"github.com/kr/pretty"

	"github.com/graphql-go/graphql"
)

func GetUser(params graphql.ResolveParams) (interface{}, error) {
	userId, err := GetUserFromGraphqlParams(params)

	if err != nil {
		return nil, err
	}
	usr, err := user.FindUser(userId)

	if err != nil {
		pretty.Log(err)
		return nil, err
	}
	pretty.Log(usr)
	return usr, nil
}

func GetPayments(params graphql.ResolveParams) (interface{}, error) {
	projectId, err := GetProjectFromGraphqlParams(params)
	if err != nil {
		return nil, err
	}
	res, err := payment.GetPaymentsByProjectId(projectId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreateProject(params graphql.ResolveParams) (interface{}, error) {
	userId, err := GetUserFromGraphqlParams(params)

	if err != nil {
		return nil, err
	}

	project, err := project.CreateProject(userId)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func GetProject(params graphql.ResolveParams) (interface{}, error) {
	//projectId, ok  := params.Args["projectId"].(string)
	projectId, err := GetProjectFromGraphqlParams(params)
	if err != nil {
		return nil, err
	}

	res, err := project.FindById(projectId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func GetUserFromGraphqlParams(params graphql.ResolveParams) (string, error) {
	return "5b8275edecbc55131a9b8733", nil
}

func GetProjectFromGraphqlParams(params graphql.ResolveParams) (string, error) {
	return "5b829ffdecbc5561ac606bf0", nil
}

func Checkout(params graphql.ResolveParams) (string, error) {
	return "", nil
}
