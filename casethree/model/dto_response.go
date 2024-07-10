package model

type DTOResponse struct {
	Id             int    `json:"Id"`
	Country        string `json:"Country"`
	CreditCardType string `json:"Credit_card_type"`
	CreditCard     string `json:"Credit_card"`
	FirstName      string `json:"First_name"`
	LastName       string `json:"Last_name"`
}
