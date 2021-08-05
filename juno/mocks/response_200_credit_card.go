package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Response200CreditCard() *MockClient {
	jsonResponse := `
		{
			"creditCardId": "string",
			"last4CardNumber": "string",
			"expirationMonth": "string",
			"expirationYear": "string"
		}
	`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	return &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 200,
				Body:       r,
			}, nil
		},
	}
}
