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

func init() {
	Client = &http.Client{}
}

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

	r, _ := http.NewRequest(http.MethodPost, urlString, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "Basic "+base64Token)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := Client.Do(r)

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

	r, _ := http.NewRequest(http.MethodPost, urlString, bytes.NewBuffer(jsonStr))
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

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

	r, _ := http.NewRequest(http.MethodGet, urlString, nil)
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

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

func (juno junoAccess) GetPlan(junoAccessAuth model.JunoAccessAuth, planID string) (*model.Plan, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/plans/" + planID

	r, _ := http.NewRequest(http.MethodGet, urlString, nil)
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		plan, err := model.FromJsonJunoPlan(resp.Body)

		if err != nil {
			return nil, err
		}

		return plan, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}

func (juno junoAccess) CreatePlan(junoAccessAuth model.JunoAccessAuth, name string, amount float64) (*model.Plan, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/plans"

	payload := map[string]interface{}{
		"name":   name,
		"amount": amount,
	}

	jsonStr, _ := json.Marshal(payload)

	r, _ := http.NewRequest(http.MethodPost, urlString, bytes.NewBuffer(jsonStr))
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		plan, err := model.FromJsonJunoPlan(resp.Body)

		if err != nil {
			return nil, err
		}

		return plan, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}

func (juno junoAccess) DisablePlan(junoAccessAuth model.JunoAccessAuth, planID string) (*model.Plan, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/plans/" + planID + "/deactivation"

	r, _ := http.NewRequest(http.MethodPost, urlString, nil)
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		plan, err := model.FromJsonJunoPlan(resp.Body)

		if err != nil {
			return nil, err
		}

		return plan, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}

func (juno junoAccess) EnablePlan(junoAccessAuth model.JunoAccessAuth, planID string) (*model.Plan, error) {
	urlString := juno.access.api + juno.access.baseUrl + "/plans/" + planID + "/activation"

	r, _ := http.NewRequest(http.MethodPost, urlString, nil)
	r.Header.Add("Authorization", "Bearer "+junoAccessAuth.AccessToken)
	r.Header.Add("X-Api-Version", "2")
	r.Header.Add("Content-Type", "application/json;charset=UTF-8")
	r.Header.Add("X-Resource-Token", juno.access.resourceToken)

	resp, err := Client.Do(r)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		plan, err := model.FromJsonJunoPlan(resp.Body)

		if err != nil {
			return nil, err
		}

		return plan, nil
	} else {
		defaultError := errors.ParseDefaultError(resp.Body)

		return nil, defaultError
	}
}
