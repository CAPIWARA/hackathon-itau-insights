package payment

import (
	"hackathon-itau-insights/server/bd"

	"github.com/mitchellh/mapstructure"
)

type Payment struct {
	Id            string `structs:"_id" json:"-" bson:"_id"`
	TransactionId string
	UserId        string
	ProjectId     string
}

func Save(transactioniD, userId, projectId string) (string, error) {
	pay := Payment{
		TransactionId: transactioniD,
		UserId:        userId,
		ProjectId:     projectId,
	}
	m := bd.MongoBuilder("payment")
	return m.Save(pay)
}

func GetPaymentsByProjectId(projectId string) ([]Payment, error) {
	var lista []Payment
	m := bd.MongoBuilder("payment")

	res, err := m.FindByField("projectId", projectId)
	if err != nil {
		return nil, err
	}
	if err := mapstructure.Decode(res, &lista); err != nil {
		return nil, err
	}
	return lista, nil
}
