package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
	"github.com/google/uuid"
)

type UserRes struct {
	Message string `json:"message"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user User

	//parse from body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//validating username
	username, _, err := db.UsernameValidation(user.Username)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if username != "" {
		http.Error(w, "username already taken", http.StatusConflict)
		return
	}

	//hahsing password
	var hash string
	if bytes, err := utils.Hash(user.Password); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	} else {
		hash = bytes
	}

	//generate user uuid
	userUUID := uuid.NewString()

	//store new user to database
	if err := db.Register(user.Username, userUUID, hash); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	//response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(UserRes{
		Message: fmt.Sprintf("User @%v registered", user.Username),
	})
}
