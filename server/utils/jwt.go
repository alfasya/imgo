package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Owner struct {
	Status   string `default:"unauthorized"`
	Username string
	UserUUID string
	UserId   int
}

var key = []byte("$hidden-dominating-crucial-burgundi!")

func CreateToken(username, uuid string, id int) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"userId":   id,
			"userUUID": uuid,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (Owner, error) {
	var owner Owner

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return owner, err
	}

	if !token.Valid {
		return owner, fmt.Errorf("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return owner, err
	}

	owner.Status = "authorized"
	owner.Username = claims["username"].(string)
	owner.UserUUID = claims["userUUID"].(string)
	owner.UserId = int(claims["userId"].(float64))

	return owner, nil
}
