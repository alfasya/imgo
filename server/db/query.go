package db

import (
	"errors"
	"fmt"

	"github.com/alfasya/imgo/utils"
)

func DbValidation(u string) (bool, error) {
	var exist bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`

	err := Pool.QueryRow(Ctx, sql, u).Scan(&exist)

	if !exist {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("Error querying: %s", err)
	}

	return true, nil
}

func PasswordValidation(u, password string) (bool, error) {
	var hash string

	sql := `SELECT hashed_password FROM users WHERE username = $1 LIMIT 1`

	err := Pool.QueryRow(Ctx, sql, u).Scan(&hash)
	if err != nil {
		return false, fmt.Errorf("Database error: %s", err)
	}

	//COMPARING IN BCRYPT
	match := utils.ComparePassword(hash, password)

	if !match {
		return false, errors.New("Password doesn't match")
	}

	return true, nil
}

func NewRegistration(u, h string) error {
	sql := `INSERT INTO users (username, hashed_password) VALUES ($1, $2)`

	_, err := Pool.Exec(Ctx, sql, u, h)
	if err != nil {
		fmt.Printf("Error registering new user int database: %s", err)
		return err
	}
	return nil
}
