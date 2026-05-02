package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

func SignJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"test": "dsajdl",
	})

	tokenString, err := token.SignedString([]byte("afj839fje9"))
	if err != nil {
		fmt.Printf("Token error: %v", err)
		return "", err
	}

	return tokenString, nil
}
