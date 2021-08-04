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

	fmt.Println(clientID)
	fmt.Println(clientSecret)
	fmt.Println(resourceToken)

	junoAccess := juno.JunoConfig().
		ClientID(clientID).
		ClientSecret(clientSecret).
		ResourceToken(resourceToken).
		Sandbox()

	junoSdk := juno.Instance(junoAccess)

	accessToken, _ := junoSdk.GetAuthToken()
	creditCard, err := junoSdk.TokenizeCard(*accessToken, "a750d6be-8940-4854-97ab-5e54a4e00716")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(creditCard)
}
