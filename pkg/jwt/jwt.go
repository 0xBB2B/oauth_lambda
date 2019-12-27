package jwt

import (
	"fmt"

	"oauth_lambda/init"

	jwt "github.com/dgrijalva/jwt-go"
)

func EncodeJWT(m jwt.MapClaims) (string, error) {
	conf := init.Conf()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, m)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(conf.HmacSecret))
	return tokenString, err
}

func DecodeJWT(s string) (jwt.MapClaims, error) {
	conf := init.Conf()
	token, err := jwt.Parse(s, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(conf.HmacSecret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
