package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	owner, ok := r.Context().Value("owner").(utils.Owner)
	if !ok {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "bad request",
		})
		return
	}

	id := owner.UserId
	uuid := owner.UserUUID
	ownersFile := r.PathValue("uuid")

	if uuid != ownersFile {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "unauthorized",
		})
	}

	filename := r.PathValue("filename")

	if err := db.DeleteQuery(id, filename); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(403)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "you are not allowed to delete the file",
		})
	}

	if err := utils.RemoveFile(owner.UserUUID, filename); err != nil {
		fmt.Printf("error deleting file: %v", err)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "internal server error",
		})
	}

	//response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "file deleted",
	})
}
