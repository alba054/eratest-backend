package main

import (
	"context"
	"database/sql"

	"casethree/helper"
	"casethree/model"

	"github.com/go-playground/validator"
)

type Service interface {
	CreateUser(ctx context.Context, request model.CreateRequest) interface{}
	FindTheMostSpentCountry(ctx context.Context) []model.DTOResponse
}

type ServiceImpl struct {
	Repository Repository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewService(repository Repository, DB *sql.DB, validate *validator.Validate) Service {
	return &ServiceImpl{
		Repository: repository,
		DB:         DB,
		Validate:   validate,
	}
}

// CreateUser implements Service.
func (service *ServiceImpl) CreateUser(ctx context.Context, request model.CreateRequest) interface{} {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	entity := model.MainEntity{
		Country:        request.Country,
		CreditCardType: request.CreditCardType,
		CreditCard:     request.CreditCard,
		FirstName:      request.FirstName,
		LastName:       request.LastName,
	}

	service.Repository.Save(ctx, tx, entity)

	return nil
}

// FindTheMostSpentCountry implements Service.
func (service *ServiceImpl) FindTheMostSpentCountry(ctx context.Context) []model.DTOResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	entities := service.Repository.FindTheMostSpentCountry(ctx, tx)

	return helper.ToEntityResponses(entities)
}
