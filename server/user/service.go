package user

import (
	"hackathon-itau-insights/server/bd"

	"github.com/mitchellh/mapstructure"
)

func FindUser(id string) (User, error) {
	var user User
	m := bd.MongoBuilder("user")
	res, err := m.FindById(id)
	if err != nil {
		return User{}, nil
	}

	if err := mapstructure.Decode(res, &user); err != nil {
		return User{}, err
	}

	return user, nil
}

var Renata = User{
	Id:    "5b823446ecbc557e49107599",
	Name:  "Renata Silva",
	Email: "renata@capiwara.com.br",
	Account: &Account{
		Agencia:        "0048",
		AgenciaDv:      "1",
		Conta:          "48800",
		ContaDv:        "1",
		BankCode:       "341",
		DocumentNumber: "46996939895",
		LegalName:      "RENATA V SILVA",
		AccountId:      17883361,
	},
	CreditCard: &CreditCard{
		CCId:       "card_cjlagrqkf00bl836dtueuoz7o",
		Number:     "4242424242424242",
		HolderName: "Renata Silva Capi",
		Exp:        "1018",
	},
	Address: &Address{
		Street:        "Rua Friedrich Von Voith",
		StreetNumber:  "808",
		Neighborhood:  "Bairro Jaragua",
		Complementary: "Apartamento 13",
		City:          "Sao Paulo",
		State:         "SP",
		Zipcode:       "02995000",
		Country:       "Brasil",
	},
	Phone: &Phone{
		Ddd:    "88",
		Ddi:    "88",
		Number: "945597088",
	},
}
