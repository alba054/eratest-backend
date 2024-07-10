package model

type CreateRequest struct {
	Country        string `validate:"required,min=1,max=100" json:"country"`
	CreditCardType string `validate:"required,min=1,max=100" json:"credit_card_type"`
	CreditCard     string `validate:"required,min=1,max=100" json:"credit_card"`
	FirstName      string `validate:"required,min=1,max=100" json:"first_name"`
	LastName       string `validate:"required,min=1,max=100" json:"last_name"`
}
