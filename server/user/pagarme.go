package user

import (
	"encoding/json"
	"net/url"

	"github.com/kr/pretty"

	"github.com/luk4z7/pagarme-go/auth"
	"github.com/luk4z7/pagarme-go/lib/bank"
	"github.com/luk4z7/pagarme-go/lib/card"
	"github.com/luk4z7/pagarme-go/lib/transaction"
)

var bankAccount bank.Account
var transactionRecord transaction.Transaction
var creditCard card.Card

func (user *User) RegisterAccount() (int, error) {
	data, _ := json.Marshal(map[string]string{
		"bank_code":       user.Account.BankCode,
		"agencia":         user.Account.Agencia,
		"agencia_dv":      user.Account.AgenciaDv,
		"conta":           user.Account.Conta,
		"conta_dv":        user.Account.ContaDv,
		"document_number": user.Account.DocumentNumber,
		"legal_name":      user.Account.LegalName,
	})

	account, err, errorsApi := bankAccount.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		pretty.Log(errorsApi)
		return 0, err
	}
	return account.Id, nil
}

func (user *User) CheckoutTransaction(amount int) (int, error) {
	data, _ := json.Marshal(map[string]interface{}{
		"amount":  amount,
		"card_id": user.CreditCard.CCId,
		"customer": map[string]interface{}{
			"name":            user.Name,
			"email":           user.Email,
			"document_number": user.Account.DocumentNumber,
			"gender":          user.Gender,
			"born_at":         user.BornAt,
			"address": map[string]string{
				"street":        user.Address.Street,
				"street_number": user.Address.StreetNumber,
				"neighborhood":  user.Address.Neighborhood,
				"complementary": user.Address.Complementary,
				"city":          user.Address.City,
				"state":         user.Address.State,
				"zipcode":       user.Address.Zipcode,
				"country":       user.Address.Country,
			},
			"phone": map[string]interface{}{
				"ddi":    user.Phone.Ddi,
				"ddd":    user.Phone.Ddd,
				"number": user.Phone.Number,
			},
		},
	})

	create, err, errosApi := transactionRecord.Create(data, url.Values{}, auth.Headers{})

	if err != nil {
		pretty.Log(errosApi)
		return 0, err
	}
	return create.Id, nil

}

func (user *User) RegisterCreditCard() (string, error) {
	data, _ := json.Marshal(map[string]string{
		"card_number":          user.CreditCard.Number,
		"card_holder_name":     user.CreditCard.HolderName,
		"card_expiration_date": user.CreditCard.Exp,
	})
	create, err, errorsApi := creditCard.Create(data, url.Values{}, auth.Headers{})
	if err != nil {
		pretty.Log(errorsApi)
		return "", err
	}
	return create.Id, nil
}
