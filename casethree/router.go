package main

import (
	"casethree/helper"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller Controller) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api", controller.GetTheMostSpentCountry)
	router.POST("/api", controller.PostUser)

	router.PanicHandler = helper.ErrorHandler

	return router
}
