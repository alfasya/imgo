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

func verifyJWT(tokenString string) error {
	token, err := jwt.Parse(tokenString, func (token *jwt.Token)) (interface{}, error) {
		return []byte("edwefwfefw"), nil
	}
	
	if err != nil {
		return err
	}

	if !token.Valid {
		fmt.Println("Invalid token: %v", err)
	}

	return nil
}
