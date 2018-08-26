package user

import (
	"hackathon-itau-insights/server/bd"
)

//Testing new design pattern
type User struct {
	Id         string `structs:"id" json:"-" bson:"id"`
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
	m := bd.MongoBuilder("user")
	id, err := m.Save(user)
	user.Id = id

	return err
}
