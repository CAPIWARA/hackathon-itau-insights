package main

import (
	"hackathon-itau-insights/server/bd"
	"hackathon-itau-insights/server/user"
	"os"

	"github.com/kr/pretty"
)

func main() {
	os.Setenv("PAGARME_APIKEY", "ak_test_7oHnHB8vfkN54xSTipWnFd5dPbuth1")

	if err := bd.ConnectDatabase(); err != nil {
		panic(err)
	}
	pretty.Log(user.Marcos.CheckoutTransaction(100))
}
