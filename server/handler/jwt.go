package handler

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("SF-JF%I9#3^Fl96i_E2")

func GenerateToken(u string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": u,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
