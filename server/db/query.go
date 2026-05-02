package db

import (
	"context"
	"fmt"
)

func Register(ctx context.Context, u string, p string) error {
	sql := `INSERT INTO users (username, hashed_password) VALUES ($1, $2)`

	_, err := Pool.Exec(ctx, sql, u, p)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return err
	}

	fmt.Printf("Created user with username: %v\n", u)
	return nil
}

func CheckUser(ctx context.Context, u string) (bool, error) {
	var exist bool

	sql := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`

	err := Pool.QueryRow(ctx, sql, u).Scan(&exist)
	if err != nil {
		fmt.Printf("Error executing query: %v", err)
		return false, err
	}

	return exist, nil
}

func Hash(ctx context.Context, u string) (string, error) {
	var hash string

	sql := `SELECT hashed_password FROM users WHERE username = $1`

	err := Pool.QueryRow(ctx, sql, u).Scan(&hash)
	if err != nil {
		fmt.Printf("Database error: %v", err)
		return "", err
	}

	return hash, nil
}
