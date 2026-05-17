package utils

import (
	"log"
	"os"
	"path/filepath"
)

func RemoveFile(userUUID, imagePath string) error {
	dst := "C:/Users/alfas/Documents/images/"

	if err := os.Remove(filepath.Join(dst, userUUID, imagePath)); err != nil {
		log.Fatal(err)
	}

	return nil
}
