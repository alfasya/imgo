package main

import (
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST api/auth/register", handlers.Register)
	mux.HandleFunc("POST api/auth/login", handlers.Login)

	db.Connect()
	defer db.Pool.Close()

	http.ListenAndServe(":8080", mux)
}
