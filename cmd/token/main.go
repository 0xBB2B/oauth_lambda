package main

import (
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	jwt "github.com/dgrijalva/jwt-go"

	jwtUtil "oauth_lambda/pkg/jwt"
)

const (
	clientID     = "1"
	clientSecret = "Ta1xNAduKyfCtZViSDDube6e0E"
)

type myEvent struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	GrantType    string `json:"grant_type"`
	Code         string `json:"code"`
}

type myResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func handleRequest(event myEvent) (myResponse, error) {
	m := jwt.MapClaims{
		"nbf":   time.Now().Unix(),
		"exp":   time.Now().AddDate(0, 0, 1).Unix(),
		"name":  "Sinon",
		"scope": "read",
	}
	jwt, err := jwtUtil.EncodeJWT(m)
	return myResponse{
		AccessToken: jwt,
		TokenType:   "bearer",
	}, err
}

func main() {
	lambda.Start(handleRequest)
}
