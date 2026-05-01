package main

import (
	"net/http"

	"github.com/alfasya/imgo/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", handlers.Register)

	http.ListenAndServe(":8080", mux)
}
