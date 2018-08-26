package main

import (
	"hackathon-itau-insights/server/bd"
	"hackathon-itau-insights/server/lgq"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/kr/pretty"

	gqlhandler "github.com/graphql-go/graphql-go-handler"
)

func main() {
	os.Setenv("PAGARME_APIKEY", "ak_test_7oHnHB8vfkN54xSTipWnFd5dPbuth1")

	if err := bd.ConnectDatabase(); err != nil {
		panic(err)
	}

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	handler := gqlhandler.New(&gqlhandler.Config{
		Schema: &lgq.Schema,
	})

	http.Handle("/graphql", handler)
	pretty.Log("starting on port :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//ctx := context.WithValue(r.Context(), "uuid", "")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")
		if r.Method == "OPTIONS" {
			log.Println("options method ")
			return
		}
		next.ServeHTTP(w, r)
	})
}
