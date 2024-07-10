package main

import (
	"net/http"

	"casethree/helper"
	"casethree/model"

	"github.com/julienschmidt/httprouter"
)

type Controller interface {
	PostUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	GetTheMostSpentCountry(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type ControllerImpl struct {
	Service Service
}

func NewController(service Service) Controller {
	return &ControllerImpl{
		Service: service,
	}
}

func (controller *ControllerImpl) PostUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CreateRequest := model.CreateRequest{}
	helper.ReadFromRequestBody(request, &CreateRequest)

	Response := controller.Service.CreateUser(request.Context(), CreateRequest)
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   Response,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ControllerImpl) GetTheMostSpentCountry(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	Responses := controller.Service.FindTheMostSpentCountry(request.Context())
	webResponse := model.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   Responses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
