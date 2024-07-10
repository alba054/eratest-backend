package helper

import "casethree/model"

func ToEntityResponse(entity model.MainEntity) model.DTOResponse {
	return model.DTOResponse{
		Id:             entity.Id,
		Country:        entity.Country,
		CreditCardType: entity.CreditCardType,
		CreditCard:     entity.CreditCard,
		FirstName:      entity.FirstName,
		LastName:       entity.LastName,
	}
}

func ToEntityResponses(entities []model.MainEntity) []model.DTOResponse {
	var responses []model.DTOResponse
	for _, entity := range entities {
		responses = append(responses, ToEntityResponse(entity))
	}
	return responses
}
