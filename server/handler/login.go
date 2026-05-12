package handler

import (
	"encoding/json"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
)

type LoginRes struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User
	//parse data from body
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//validating username
	username, id, err := db.UsernameValidation(user.Username)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if username == "" {
		http.Error(w, "username or password is wrong", http.StatusUnauthorized)
		return
	}

	//validating password
	match, err := db.PasswordValidation(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	if !match {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//create token
	tokenString, err := utils.CreateToken(user.Username, id)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	//Response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(LoginRes{
		Message: "login success",
		Token:   tokenString,
	})
}
