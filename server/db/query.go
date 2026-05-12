package db

import (
	"fmt"

	"github.com/alfasya/imgo/utils"
)

func UploadQuery(filename string, size int, path string) error {
	query := `INSERT INTO images (name, size, path) VALUES ($1, $2, $3)`
	_, err := Pool.Exec(Ctx, query, filename, size, path)
	if err != nil {
		fmt.Printf("Error executing database: %v", err)
		return err
	}

	return nil
}

func Register(username, hash string) error {
	query := `INSERT INTO users (username, hashed_password) VALUES ($1, $2)`

	_, err := Pool.Exec(Ctx, query, username, hash)
	if err != nil {
		fmt.Printf("Error executing database: %v", err)
		return err
	}

	return nil
}

func PasswordValidation(username, password string) (bool, error) {
	var hash string
	query := `SELECT (hashed_password) FROM users WHERE username = $1`

	if err := Pool.QueryRow(Ctx, query, username).Scan(&hash); err != nil {
		fmt.Printf("Error executing database: %v", err)
		return false, err
	}

	match := utils.ComparePassword(password, hash)
	if !match {
		return false, nil
	}

	return true, nil
}

func UsernameValidation(username string) (bool, error) {
	var user string
	query := `SELECT username FROM users WHERE username = $1`

	err := Pool.QueryRow(Ctx, query, username).Scan(&user)
	if err != nil {
		fmt.Printf("Error querying table: %v", err)
		return false, nil
	}

	return true, nil
}
