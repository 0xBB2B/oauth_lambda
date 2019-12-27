package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

func main() {
	r := gin.Default()
	r.GET("/callback", callback)

	r.Run(":8002")
}

func callback(c *gin.Context) {
	code := c.Query("code")
	tokenURL := fmt.Sprintf("http://127.0.0.1:8000/oauth/token?client_id=1&redirect_uri=http://127.0.0.1:8002/callback&client_secret=Ta1xNAduKyfCtZViSDDube6e0E&grant_type=authorization_code&code=%s", code)
	jsonStr := httpGet(tokenURL)
	var t token
	json.Unmarshal([]byte(jsonStr), &t)
	c.JSON(http.StatusOK, gin.H{
		"data": httpClient(t),
	})
}

func httpClient(t token) string {
	client := &http.Client{}
	url := "http://127.0.0.1:8001/data"
	reqest, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
	reqest.Header.Add("Authorization", "Bearer "+t.AccessToken)
	resp, err := client.Do(reqest)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return string(body)
}

func httpGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}
	return string(body)
}
