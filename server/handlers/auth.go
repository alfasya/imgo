package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/utils"
)

type Res struct {
	Message string `json:"message"`
}

type NewUserRes struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if newUser.Username == "" || newUser.Password == "" {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		return
	}

	exist, err := db.CheckUser(r.Context(), newUser.Username)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	if exist {
		http.Error(w, "username already exists.", http.StatusConflict)
		return
	}

	hashedPassword, err := utils.HashPassword(newUser.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
	}

	if err := db.Register(r.Context(), newUser.Username, hashedPassword); err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	if err := json.NewEncoder(w).Encode(NewUserRes{
		Message: "User created.",
		Data:    newUser,
	}); err != nil {
		log.Printf("Internal server error: %v", err)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	exist, err := db.CheckUser(r.Context(), user.Username)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	if !exist {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	hash, err := db.Hash(r.Context(), user.Username)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	match := utils.CheckPasswordHash(user.Password, hash)

	if !match {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	tokenString, err := utils.SignJWT()
	if err != nil {
		fmt.Printf("Error signing JWT: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(Res{
		Message: tokenString,
	})
}
