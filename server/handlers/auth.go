package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)

	if err := json.NewEncoder(w).Encode(NewUserRes{
		Message: "User created.",
		Data:    newUser,
	}); err != nil {
		log.Printf("Internal server error: %v", err)
	}
}
