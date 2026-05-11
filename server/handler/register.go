package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
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

	//store new user to database
	if err := db.Register(user.Username, user.Password); err != nil {
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
