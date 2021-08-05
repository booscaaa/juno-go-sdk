package model

import (
	"encoding/json"
	"fmt"
	"io"
)

type CreditCard struct {
	CreditCardId    string `json:"creditCardId"`
	Last4CardNumber string `json:"last4CardNumber"`
	ExpirationMonth string `json:"expirationMonth"`
	ExpirationYear  string `json:"expirationYear"`
}

func (creditCard CreditCard) isValid() (bool, error) {
	if creditCard.CreditCardId == "" {
		return false, fmt.Errorf("CreditCardId not be empty")
	}

	if creditCard.Last4CardNumber == "" {
		return false, fmt.Errorf("Last4CardNumber not be empty")
	}

	if creditCard.ExpirationMonth == "" {
		return false, fmt.Errorf("ExpirationMonth not be empty")
	}

	if creditCard.ExpirationYear == "" {
		return false, fmt.Errorf("ExpirationYear not be empty")
	}

	return true, nil
}

func FromJsonJunoCreditCard(body io.ReadCloser) (*CreditCard, error) {
	creditCard := CreditCard{}
	json.NewDecoder(body).Decode(&creditCard)

	if isValid, err := creditCard.isValid(); !isValid {
		return nil, err
	}

	return &creditCard, nil
}
