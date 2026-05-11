package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

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
		dst, err := os.Create(filepath.Join("C:/Users/alfas/Documents", newFilename))
		if err != nil {
			fmt.Errorf("Error creating filepath: %v", err)
			return
		}

		//Reading and writing file to the disk
		byteFile, err := file.Open()
		if err != nil {
			fmt.Errorf("Error reading file: %v", err)
			return
		}
		io.Copy(dst, byteFile)

		//Close streaming
		dst.Close()
		byteFile.Close()
	}

	//Create file

	//Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(Res{
		Message: "OK",
	})
}
