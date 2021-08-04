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
	clientID := viper.GetString("juno.client_id")
	clientSecret := viper.GetString("juno.client_secret")
	resourceToken := viper.GetString("juno.resource_token")

	junoAccess := juno.JunoConfig().
		ClientID(clientID).
		ClientSecret(clientSecret).
		ResourceToken(resourceToken).
		Sandbox()

	junoSdk := juno.Instance(junoAccess)

	accessToken, _ := junoSdk.GetAuthToken()
	// creditCard, err := junoSdk.TokenizeCard(*accessToken, "a750d6be-8940-4854-97ab-5e54a4e00716")

	plans, err := junoSdk.GetPlans(*accessToken)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(plans)
}
