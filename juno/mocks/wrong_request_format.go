package mocks

import (
	"fmt"
	"net/http"
)

func WrongRequestFormat() *MockClient {
	return &MockClient{
		MockDo: func(*http.Request) (*http.Response, error) {
			return nil, fmt.Errorf("mock error return")
		},
	}
}
