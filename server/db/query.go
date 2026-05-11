package db

import (
	"fmt"
)

func UploadQuery(filename string, size int, path string) error {
	sql := `INSERT INTO images (name, size, path) VALUES ($1, $2, $3)`
	_, err := Pool.Exec(Ctx, sql, filename, size, path)
	if err != nil {
		fmt.Printf("Error executing database: %v", err)
		return err
	}

	return nil
}
