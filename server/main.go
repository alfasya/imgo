package main

import (
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", handlers.Register)

	db.Connect()

	http.ListenAndServe(":8080", mux)
}
