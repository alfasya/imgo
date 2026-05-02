package db

import (
	"fmt"
)

func Register(u string, p string) error {
	sql := `INSERT INTO users VALUES ($1, $2)`

	_, err := Pool.Exec(ctx, sql, u, p)
	if err != nil {
		return fmt.Errorf("Failed to execute query: %w", err)
	}

	fmt.Printf("Created user with username: %v\n", u)
	return nil
}
