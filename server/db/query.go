package db

import (
	"errors"
	"fmt"

	"github.com/alfasya/imgo/models"
	"github.com/alfasya/imgo/utils"
	"github.com/jackc/pgx/v5"
)

func UploadQuery(filename string, size int, path string, id int) error {
	query := `INSERT INTO images (source_name, source_size, source_path, user_id) VALUES ($1, $2, $3, $4)`
	_, err := Pool.Exec(Ctx, query, filename, size, path, id)
	if err != nil {
		fmt.Printf("Error executing database: %v", err)
		return err
	}

	return nil
}

func DeleteQuery(id int, filename string) error {
	query := `DELETE FROM images WHERE source_name = $1 AND user_id = $2`

	_, err := Pool.Exec(Ctx, query, filename, id)
	if err != nil {
		fmt.Printf("Error querying table: %v", err)
		return err
	}

	return nil
}

func Register(username, uuid, hash string) error {
	query := `INSERT INTO users (username, uuid, hashed_password) VALUES ($1, $2, $3)`

	_, err := Pool.Exec(Ctx, query, username, uuid, hash)
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

func UsernameValidation(username string) (string, string, int, error) {
	var u string
	var uuid string
	var id int
	query := `SELECT username, uuid, id FROM users WHERE username = $1`

	err := Pool.QueryRow(Ctx, query, username).Scan(&u, &uuid, &id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", "", 0, nil
		}
		fmt.Printf("Error querying table: %v", err)
		return "", "", 0, err
	}

	return u, uuid, id, nil
}

func GetImages(id int) ([]models.Image, error) {
	query := `SELECT source_name, source_size, source_path FROM images WHERE user_id = $1`

	rows, err := Pool.Query(Ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("error querying rows: %v", err)
	}

	defer rows.Close()

	var images []models.Image
	for rows.Next() {
		var image models.Image
		err := rows.Scan(
			&image.Name,
			&image.Size,
			&image.Path,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning rows: %v", err)
		}
		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iteratinng rows: %v", err)
	}

	return images, nil
}
