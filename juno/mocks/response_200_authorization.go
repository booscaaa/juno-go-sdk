package mocks

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func Respose200Authorization() *MockClient {
	jsonResponse := `
	  {
		"access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXvCJ9.eyJ1c2VyX25hbWUiOiJ1c2VyQGRvbWFpbi5jb20iLCJzY29wZSI6WyJhbGwiXSwiZXhwIjoxNTczMDcsMDUyLCJqdGkiOiI5NzNhMGY1Yi01NzBmLTQ5OTMtYmViZi1iNTZmZGQyMjgwNGQiLCJjbGllbnRfaWQiOialeGVtcGxvLWNsaWVudC1pZCJ9.aI8b3TqmoXKN9JZG27t64z6GBM7u9BmcDoELvBmXOAnT_g-NZ3qo83ptpTtZXHGOzx7J0hFdzVZ60feCN7gBfM-UhOwKaRwKJwmA41ZPqNyInoRBieuLCrBua8DwNHJ-BroQBWvXgeuFmyukqz1nRGoJTMuov6gbRjThdsYM0pBS4bsLM9QhZEnPkG1C7PSSNBYpvsVdBSvSMWZk8ulrKennMn305O7WPL25W4HyU3RorJK2AyiQ2pYFIDf--QdpFjzALYpwj_WKi0W2noBUKM4lCnHgsgM6ezxO56n68eB5PQzDLtI3CqRgwMgBcys5aLUWO5Y71AjC1rerEOLOOQ",
		"token_type": "bearer",
		"expires_in": 86399,
		"scope": "all",
		"user_name": "user@domain.com",
		"jti": "973a0f5b-570f-4993-bebf-b56fdd22804d"
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
