package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Response200GetPlan() *MockClient {
	jsonResponse := `
		{
			"id": "pln_D539CC5AF0E87FB1",
			"createdOn": "2020-06-22 07:22:18",
			"name": "Primeiro plano",
			"frequency": "MONTHLY",
			"status": "ACTIVE",
			"amount": 100.01,
			"_links": [
				{
				"href": "https://{url_resource_server}/plans/pln_D539CC5AF0E87FB1"
				}
			]
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
