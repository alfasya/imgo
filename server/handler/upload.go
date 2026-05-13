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

type Res struct {
	Message string `json:"message"`
}

func Upload(w http.ResponseWriter, r *http.Request) {
	//Parse form
	r.ParseMultipartForm(10 << 20)
	files := r.MultipartForm.File["images"]

	for _, file := range files {
		//Create destination
		ext := filepath.Ext(file.Filename)
		newFilename := uuid.NewString() + ext
		path := filepath.Join("C:/Users/alfas/Documents", newFilename)
		dst, err := os.Create(path)
		if err != nil {
			fmt.Printf("Error creating filepath: %v", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}

		//Reading and writing file to the disk
		byteFile, err := file.Open()
		if err != nil {
			fmt.Printf("Error reading file: %v", err)
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		io.Copy(dst, byteFile)

		owner := r.Context().Value("owner").(utils.Owner)

		//Add file's metadata to the database
		if err := db.UploadQuery(newFilename, int(file.Size), path, owner.UserId); err != nil {
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		//Close streaming
		dst.Close()
		byteFile.Close()
	}

	//Creating file

	//Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Res{
		Message: "OK",
	})
}
