package model

import (
	"encoding/json"
	"fmt"
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

func (junoAccessToken JunoAccessAuth) isValid() (bool, error) {
	if junoAccessToken.AccessToken == "" {
		return false, fmt.Errorf("AccessToken not be empty")
	}

	if junoAccessToken.TokenType == "" {
		return false, fmt.Errorf("TokenType not be empty")
	}

	if junoAccessToken.ExpiresIn == 0 {
		return false, fmt.Errorf("ExpiresIn not be empty")
	}

	if junoAccessToken.Scope == "" {
		return false, fmt.Errorf("Scope not be empty")
	}

	if junoAccessToken.UserName == "" {
		return false, fmt.Errorf("UserName not be empty")
	}

	if junoAccessToken.Jti == "" {
		return false, fmt.Errorf("Jti not be empty")
	}

	return true, nil
}

func FromJsonJunoAccessAuth(body io.ReadCloser) (*JunoAccessAuth, error) {
	junoAccessAuth := JunoAccessAuth{}
	json.NewDecoder(body).Decode(&junoAccessAuth)

	if isValid, err := junoAccessAuth.isValid(); !isValid {
		return nil, err
	}

	return &junoAccessAuth, nil
}
