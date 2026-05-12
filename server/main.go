package main

import (
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handler.Register)
	mux.HandleFunc("POST /login", handler.Login)
	mux.HandleFunc("POST /upload", handler.Upload)
	mux.HandleFunc("GET /gallery", handler.Gallery)

	db.Connect()
	defer db.Pool.Close()

	fmt.Println("Server is running and listening at port 8080...")
	http.ListenAndServe(":8080", mux)
}
