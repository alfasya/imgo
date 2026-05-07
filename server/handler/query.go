package handler

import (
	"errors"
	"fmt"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
	"github.com/jackc/pgx/v5"
)

func DbValidation(u string) (bool, error) {
	var exist bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`

	err := db.Pool.QueryRow(db.Ctx, sql, u).Scan(&exist)

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

	sql := `SELECT hashed_password FROM users WHERE username = $1 LIMIT 1;`

	err := db.Pool.QueryRow(db.Ctx, sql, u).Scan(&hash)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil
	}

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
