package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"

	rs "oauth_lambda/pkg/randstring"
)

const (
	clientID    = "1"
	redirectURI = "https://127.0.0.1/callback"
)

type myEvent struct {
	UUID        string `json:"uuid"`
	ClientID    string `json:"client_id"`
	RedirectURI string `json:"redirect_uri"`
}

type myResponse struct {
	URI string `json:"calback"`
}

func handleRequest(event myEvent) (myResponse, error) {

	// TODO: Login view

	code := rs.RandString(16)
	return myResponse{URI: fmt.Sprintf("%s?code=%s", event.RedirectURI, code)}, nil
}

func main() {
	lambda.Start(handleRequest)
}
