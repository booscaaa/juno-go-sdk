package juno

import "github.com/booscaaa/juno-go-sdk/juno/model"

const (
	PRODUCTION          = "https://api.juno.com.br"
	SANDBOX             = "https://sandbox.boletobancario.com"
	BASE_URL_PRODUCTION = ""
	BASE_URL_SANDBOX    = "/api-integration"
)

type JunoAccess struct {
	api           string
	baseUrl       string
	clientID      string
	clientSecret  string
	resourceToken string
}

func JunoConfig() JunoAccess {
	return JunoAccess{
		api:     PRODUCTION,
		baseUrl: BASE_URL_PRODUCTION,
	}
}

func (junoAccess JunoAccess) ClientID(clientID string) JunoAccess {
	junoAccess.clientID = clientID

	return junoAccess
}

func (junoAccess JunoAccess) ClientSecret(clientSecret string) JunoAccess {
	junoAccess.clientSecret = clientSecret

	return junoAccess
}

func (junoAccess JunoAccess) ResourceToken(resourceToken string) JunoAccess {
	junoAccess.resourceToken = resourceToken

	return junoAccess
}

func (junoAccess JunoAccess) Sandbox() JunoAccess {
	junoAccess.api = SANDBOX
	junoAccess.baseUrl = BASE_URL_SANDBOX

	return junoAccess
}

type JunoAccessRepository interface {
	GetAuthToken() (*model.JunoAccessAuth, error)
	TokenizeCard(junoAccessAuth model.JunoAccessAuth, creditCardHash string) (*model.CreditCard, error)
}
