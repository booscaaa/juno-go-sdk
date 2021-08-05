package juno

import (
	"testing"

	"github.com/booscaaa/juno-go-sdk/juno/mocks"
	"github.com/smartystreets/assertions/assert"
	"github.com/smartystreets/assertions/should"
)

func TestSholdBeReturnJunoConfig(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	assert.So(junoAccess.api, should.Equal, "https://sandbox.boletobancario.com").Fatal()
	assert.So(junoAccess.clientID, should.Equal, "123").Fatal()
	assert.So(junoAccess.clientSecret, should.Equal, "1234").Fatal()
	assert.So(junoAccess.resourceToken, should.Equal, "12345abc").Fatal()
}

func TestSholdBeReturnJunoInstance(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	assert.So(junoSdk, should.NotBeEmpty).Fatal()
	assert.So(junoSdk, should.NotBeNil).Fatal()
}

func TestSholdBeReturnJunoAuthToken(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.So(authToken.AccessToken, should.NotBeEmpty).Fatal()
}

func TestSholdBeReturnJunoDefaultErrorOnGetAuthToken(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.ResponseDefaultError()

	authToken, err := junoSdk.GetAuthToken()

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(authToken, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnAnyKeyEmptyOnGetAuthToken(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Response200WithoutKey()

	authToken, err := junoSdk.GetAuthToken()

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(authToken, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnRequestNotFormated(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.WrongRequestFormat()

	authToken, err := junoSdk.GetAuthToken()

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(authToken, should.BeNil).Fatal()
}

func TestSholdBeReturnJunoCreditCard(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.Response200CreditCard()

	creditCard, err := junoSdk.TokenizeCard(*authToken, "")

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.So(creditCard.CreditCardId, should.NotBeEmpty).Fatal()
}

func TestSholdBeReturnJunoDefaultErrorOnGetCreditCard(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.ResponseDefaultError()

	creditCard, err := junoSdk.TokenizeCard(*authToken, "")

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(creditCard, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnAnyKeyEmptyOnGetCreditCard(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.Response200WithoutKey()

	creditCard, err := junoSdk.TokenizeCard(*authToken, "")

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(creditCard, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnRequestNotFormatedOnGetCreditCard(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.WrongRequestFormat()

	creditCard, err := junoSdk.TokenizeCard(*authToken, "")

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(creditCard, should.BeNil).Fatal()
}

func TestSholdBeReturnJunoSlicePlans(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.Response200GetPlans()

	plans, err := junoSdk.GetPlans(*authToken)

	if err != nil {
		t.Errorf(err.Error())
	}

	assert.So(plans, should.NotBeEmpty).Fatal()
	assert.So(plans, should.NotBeNil).Fatal()
	assert.So(len(*plans), should.Equal, 1).Fatal()
}

func TestSholdBeReturnJunoDefaultErrorOnGetPlans(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.ResponseDefaultError()

	plans, err := junoSdk.GetPlans(*authToken)

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(plans, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnAnyKeyEmptyOnGetPlans(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.Response200WithoutKey()

	plans, err := junoSdk.GetPlans(*authToken)

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(plans, should.BeNil).Fatal()
}

func TestSholdBeReturnErrorOnRequestNotFormatedOnGetPlans(t *testing.T) {
	junoAccess := JunoConfig().
		ClientID("123").
		ClientSecret("1234").
		ResourceToken("12345abc").
		Sandbox()

	junoSdk := Instance(junoAccess)

	Client = mocks.Respose200Authorization()

	authToken, err := junoSdk.GetAuthToken()

	Client = mocks.WrongRequestFormat()

	plans, err := junoSdk.GetPlans(*authToken)

	if err == nil {
		t.Errorf("Shold be return error, not ok")
	}

	assert.So(plans, should.BeNil).Fatal()
}
