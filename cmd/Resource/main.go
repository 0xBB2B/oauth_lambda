package main

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	hmacSecret = "1885df74d00dbbe19274c6d955feeb5b"
)

func main() {
	r := gin.Default()
	authorized := r.Group("/")
	authorized.Use(authRequired())
	{
		authorized.GET("/data", data)
	}

	r.Run(":8001")
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		s := strings.Split(auth, " ")
		if len(s) != 2 || s[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authorization error",
			})
			c.Abort()
			return
		}

		jwt, err := decodeJWT(s[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": fmt.Sprintf("%s", err),
			})
			c.Abort()
			return
		}
		c.Set("name", jwt["name"].(string))
	}
}

func data(c *gin.Context) {
	name := c.MustGet("name").(string)
	c.JSON(http.StatusOK, gin.H{
		"name": name,
	})
}

func decodeJWT(s string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(hmacSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
