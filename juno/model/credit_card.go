package model

import (
	"encoding/json"
	"io"
)

type CreditCard struct {
	CreditCardId    string `json:"creditCardId"`
	Last4CardNumber string `json:"last4CardNumber"`
	ExpirationMonth int    `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
}

func FromJsonJunoCreditCard(body io.ReadCloser) (*CreditCard, error) {
	creditCard := CreditCard{}
	err := json.NewDecoder(body).Decode(&creditCard)

	if err != nil {
		return nil, err
	}

	return &creditCard, nil
}
