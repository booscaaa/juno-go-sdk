package juno

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/booscaaa/juno-go-sdk/juno/errors"
	"github.com/booscaaa/juno-go-sdk/juno/model"
)

type junoAccess struct {
	access JunoAccess
}

func Instance(access JunoAccess) JunoAccessRepository {
	return &junoAccess{
		access: access,
	}
}

func (juno junoAccess) GetAuthToken() (*model.JunoAccessAuth, error) {
	urlString := juno.access.api + "/authorization-server/oauth/token"

	base64Token := base64.StdEncoding.EncodeToString([]byte(juno.access.clientID + ":" + juno.access.clientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlString, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "Basic "+base64Token)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		junoAccessAuth, err := model.FromJsonJunoAccessAuth(resp.Body)

		if err != nil {
			return nil, err
		}

		return junoAccessAuth, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}

func (juno junoAccess) TokenizeCard(junoAccessAuth model.JunoAccessAuth, creditCardHash string) (*model.CreditCard, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/credit-cards/tokenization"

	payload := map[string]interface{}{
		"creditCardHash": creditCardHash,
	}

	jsonStr, _ := json.Marshal(payload)

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlString, bytes.NewBuffer(jsonStr))
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		junoCreditCard, err := model.FromJsonJunoCreditCard(resp.Body)

		if err != nil {
			return nil, err
		}

		return junoCreditCard, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}

func (juno junoAccess) GetPlans(junoAccessAuth model.JunoAccessAuth) (*[]model.Plan, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/plans"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodGet, urlString, nil)
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		plans, err := model.FromJsonJunoPlans(resp.Body)

		if err != nil {
			return nil, err
		}

		return plans, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}
