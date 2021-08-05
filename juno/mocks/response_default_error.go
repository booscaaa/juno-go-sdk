package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func ResponseDefaultError() *MockClient {
	jsonResponse := `
	  {
		"timestamp": "string",
		"status": 0,
		"error": "string",
		"details": [
		  {
			"field": "optional",
			"message": "string",
			"errorCode": "string"
		  }
		],
		"path": "string"
	  }
	`
	// create a new reader with that JSON
	r := ioutil.NopCloser(bytes.NewReader([]byte(jsonResponse)))
	return &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: 400,
				Body:       r,
			}, nil
		},
	}
}
