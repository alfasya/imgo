package main

import (
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /upload", handler.Upload)

	fmt.Println("Server is running...")
	http.ListenAndServe(":8080", mux)
}
