package payment

import "hackathon-itau-insights/server/bd"

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
