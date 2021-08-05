<p align="center">
  <h1 align="center">Juno API - Golang SDK</h1>
  <p align="center">Juno API (Open Banking) (2.0.0) </p>
  <p align="center">
    <a href="https://github.com/booscaaa/juno-go-sdk/releases/latest"><img alt="Release" src="https://img.shields.io/github/v/release/booscaaa/juno-go-sdk.svg?style=for-the-badge"></a>
    <a href="/LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-red.svg?style=for-the-badge"></a>
    <a href="https://github.com/booscaaa/juno-go-sdk/actions/workflows/test.yaml"><img alt="Build status" src="https://img.shields.io/github/workflow/status/booscaaa/juno-go-sdk/Test?style=for-the-badge"></a>
    <a href="https://codecov.io/gh/booscaaa/juno-go-sdk"><img alt="Coverage" src="https://img.shields.io/codecov/c/github/booscaaa/juno-go-sdk/master.svg?style=for-the-badge"></a>
  </p>
</p>

<br>

## Why?

This project is part of my personal portfolio, so, I'll be happy if you could provide me any feedback about the project, code, structure or anything that you can report that could make me a better developer!

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/).

<br>

## Functionalities

- Get authentication token
- Create tokenized card
- Get plans for signatures

<br>

## Getting Started

### Prerequisites

To run this project in the development mode, you'll need to have a basic environment to run:

- A Golang SDK, that can be found [here](https://golang.org/).

<br>

### Installing

```bash
$ go get juno-go-sdk
```

<br>

### Create a config.json file inside your project like this
**The access client_id, client_secret and resource_token can be found into Juno account**
```json
{
  "juno": {
    "client_id": "cliend id",
    "client_secret": "client secret",
    "resource_token": "resource token"
  }
}
```
<br>
<br>

### Create main.go file
```go
package main

import (
    "fmt"

    "juno-go-sdk/juno"
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

    //Get list plans from juno api
    plans, err := junoSdk.GetPlans(*accessToken)

    if err != nil {
        fmt.Println(err)
    }

    //see more examples into examples folder
}
```

# Running local

```bash
go run main.go
```


<br>
<br>
<br>


## Api application built With

- [Viper](https://github.com/spf13/viper)

<br>
<br>
<br>

## Contributing

You can send how many PR's do you want, I'll be glad to analyze and accept them! And if you have any question about the project...

Email-me: boscardinvinicius@gmail.com

Connect with me at [LinkedIn](https://www.linkedin.com/in/booscaaa/)

Thank you!

## License

This project is licensed under the MIT License - see the [LICENSE.md](juno-go-sdk/blob/master/LICENSE) file for details