package main

import (
	"fmt"

	"github.com/booscaaa/juno-go-sdk/juno"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	// Get configuration variavables
	clientID := viper.GetString("juno.client_id")
	clientSecret := viper.GetString("juno.client_secret")
	resourceToken := viper.GetString("juno.resource_token")

	// Configure Sandbox access into Juno api
	junoAccess := juno.JunoConfig().
		ClientID(clientID).
		ClientSecret(clientSecret).
		ResourceToken(resourceToken).
		Sandbox()

	junoSdk := juno.Instance(junoAccess)

	// Get authentication token from juno api
	accessToken, err := junoSdk.GetAuthToken()

	if err != nil {
		fmt.Println(err)
	}

	// Tokenize credit card by hash. How to crypt hash can be found hear: https://dev.juno.com.br/api/v2#tag/Obtendo-o-hash
	creditCard, err := junoSdk.TokenizeCard(*accessToken, "e6bfca8c-ef97-4707-bc2a-e95ac26898be")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(creditCard)

	//Get list plans from juno api
	plans, err := junoSdk.GetPlans(*accessToken)

	if err != nil {
		fmt.Println(err)
	}

	for _, plan := range *plans {
		//Get plan by id from juno api
		newPlan, err := junoSdk.GetPlan(*accessToken, plan.ID)

		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(newPlan)
	}
}
