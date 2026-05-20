package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
	"github.com/google/uuid"
)

func Response(w http.ResponseWriter, statusCode int, message string) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{
		"message": message,
	})
	return nil
}

func Upload(w http.ResponseWriter, r *http.Request) {
	owner, ok := r.Context().Value("owner").(utils.Owner)
	if !ok {
		Response(w, 400, "invalid type")
		return
	}
	//Parse form
	r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["images"]

	dir := filepath.Join("C:/Users/alfas/Documents/images", owner.UserUUID)

	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.Mkdir(dir, 0755)
	}

	for _, file := range files {
		//Create destination
		ext := filepath.Ext(file.Filename)
		newFilename := uuid.NewString() + ext
		path := filepath.Join(dir, newFilename)
		dst, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating filepath: %v", err)
			Response(w, 500, "internal server error")
			return
		}

		//Reading and writing file to the disk
		byteFile, err := file.Open()
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			Response(w, 500, "internal server error")
			return
		}
		io.Copy(dst, byteFile)

		//Add file's metadata to the database
		if err := db.UploadQuery(newFilename, int(file.Size), path, owner.UserId); err != nil {
			Response(w, 500, "internal server error")
			return
		}

		//Close streaming
		dst.Close()
		byteFile.Close()
	}

	Response(w, 200, fmt.Sprintf("Uploaded %d images", len(files)))
}
