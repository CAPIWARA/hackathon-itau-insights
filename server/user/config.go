package user

import (
	"hackathon-itau-insights/server/bd"
)

var db = bd.MongoBuilder("user")

//Testing new design pattern
type User struct {
	Id         string `structs:"_id" json:"-" bson:"_id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Gender     string `json:"gender"`
	BornAt     string `json:"bornAt"`
	Account    *Account
	CreditCard *CreditCard
	Address    *Address
	Phone      *Phone
}

type Account struct {
	BankCode       string
	Agencia        string
	AgenciaDv      string
	Conta          string
	ContaDv        string
	DocumentNumber string
	LegalName      string
	AccountId      int
}

type CreditCard struct {
	Number     string
	HolderName string
	CCId       string
	Exp        string
}

type Address struct {
	Street        string
	StreetNumber  string
	Neighborhood  string
	Complementary string
	City          string
	State         string
	Zipcode       string
	Country       string
}

type Phone struct {
	Ddi    string
	Ddd    string
	Number string
}

func (user *User) Save() error {
	id, err := db.Save(user)
	user.Id = id

	return err
}

var Marcos = User{
	Name:  "Marcos Vinicius",
	Email: "marcos@capiwara.com.br",
	Account: &Account{
		Agencia:        "0048",
		AgenciaDv:      "1",
		Conta:          "48800",
		ContaDv:        "1",
		BankCode:       "341",
		DocumentNumber: "46996939895",
		LegalName:      "MARCOS V CARDOSO",
		AccountId:      17883361,
	},
	CreditCard: &CreditCard{
		CCId:       "card_cjlagrqkf00bl836dtueuoz7o",
		Number:     "4242424242424242",
		HolderName: "Marcos Vinicius Capi",
		Exp:        "1018",
	},
	Address: &Address{
		Street:        "Rua de exemplo",
		StreetNumber:  "808",
		Neighborhood:  "Bairro de exemplo",
		Complementary: "Apartamento 8",
		City:          "Cidade",
		State:         "Lordaeron",
		Zipcode:       "47850000",
		Country:       "Lordaeron",
	},
	Phone: &Phone{
		Ddd:    "88",
		Ddi:    "88",
		Number: "945597088",
	},
}
