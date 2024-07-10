package main

import (
	"net/http"

	"casethree/helper"

	"github.com/go-playground/validator"
)

func main() {
	db := helper.NewDB()
	validate := validator.New()
	repository := NewRepository()
	service := NewService(repository, db, validate)
	controller := NewController(service)
	router := NewRouter(controller)
	LoadCSVToDatabase(db, "userdata.csv")
	LoadTransactionDatabase(db, "transactiondata.csv")

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
