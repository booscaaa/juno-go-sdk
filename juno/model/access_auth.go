package model

import (
	"encoding/json"
	"io"
)

type JunoAccessAuth struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	UserName    string `json:"user_name"`
	Jti         string `json:"jti"`
}

func FromJsonJunoAccessAuth(body io.ReadCloser) (*JunoAccessAuth, error) {
	junoAccessAuth := JunoAccessAuth{}
	err := json.NewDecoder(body).Decode(&junoAccessAuth)

	if err != nil {
		return nil, err
	}

	return &junoAccessAuth, nil
}
