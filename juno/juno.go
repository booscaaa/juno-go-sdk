package juno

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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
	base64Token := base64.StdEncoding.EncodeToString([]byte(juno.access.clientID + ":" + juno.access.clientSecret))

	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	urlString := juno.access.api + "/authorization-server/oauth/token"

	client := &http.Client{}
	r, _ := http.NewRequest(http.MethodPost, urlString, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "Basic "+base64Token)
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)

	if resp.StatusCode == 200 {
		junoAccessAuth, err := model.FromJsonJunoAccessAuth(resp.Body)

		if err != nil {
			return nil, err
		}

		return junoAccessAuth, nil
	}

	return nil, fmt.Errorf("Request failed")
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

	resp, _ := client.Do(r)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	if resp.StatusCode == 200 {
		junoCreditCard, err := model.FromJsonJunoCreditCard(resp.Body)

		if err != nil {
			return nil, err
		}

		return junoCreditCard, nil
	}

	return nil, fmt.Errorf("Request failed")
}
