package main

import (
	"encoding/json"
	"errors"

	"github.com/aws/aws-lambda-go/lambda"
	uuid "github.com/satori/go.uuid"
)

const (
	clientID    = "1"
	redirectURI = "https://127.0.0.1/callback"
)

type myEvent struct {
	UUID         string `json:"uuid"`
	ClientID     string `json:"client_id"`
	RedirectURI  string `json:"redirect_uri"`
	ResponseType string `json:"response_type"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"` // refresh
}

type myResponse struct {
	UUID string `json:"uuid"`
}

// ErrJSON type
type ErrJSON map[string]interface{}

// JSONString is ErrJson type to json string
func (ej ErrJSON) JSONString() string {
	j, _ := json.Marshal(ej)
	return string(j)
}

func handleRequest(event myEvent) (interface{}, error) {
	// TODO: DB check
	if event.ClientID != clientID {
		ej := ErrJSON{
			"status":  401,
			"message": "Chient ID error",
		}
		return nil, errors.New(ej.JSONString())
	}
	if event.RedirectURI != redirectURI {
		ej := ErrJSON{
			"status":  400,
			"message": "Redirect URI error",
		}
		return nil, errors.New(ej.JSONString())
	}
	if event.ResponseType != "authorization_code" {
		ej := ErrJSON{
			"status":  400,
			"message": "Response type error",
		}
		return nil, errors.New(ej.JSONString())
	}

	uu := uuid.NewV4()
	event.UUID = uu.String()
	// TODO: DB save event

	return myResponse{UUID: event.UUID}, nil
}

func main() {
	lambda.Start(handleRequest)
}
