package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte("$hidden-dominating-crucial-burgundi!")

func CreateToken(username string) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"duration": "forever",
		})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	owner := claims["username"].(string)

	return owner, nil
}
