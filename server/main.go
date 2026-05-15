package main

import (
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/handler"
	"github.com/alfasya/imgo/middlewares"
	"github.com/rs/cors"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handler.Register)
	mux.HandleFunc("POST /login", handler.Login)
	mux.Handle("POST /upload", middlewares.Auth(http.HandlerFunc(handler.Upload)))
	mux.Handle("GET /gallery", middlewares.Auth(http.HandlerFunc(handler.Gallery)))
	mux.Handle("DELETE /images/{uuid}/{filename}", middlewares.Auth(http.HandlerFunc(handler.Delete)))

	mux.Handle("GET /images/{uuid}/", handler.FileServer())

	db.Connect()
	defer db.Pool.Close()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"http://127.0.0.1:5500",
			"http://localhost:5500",
		},
		AllowedMethods: []string{
			"GET", "POST", "DELETE", "OPTIONS",
		},
		AllowedHeaders: []string{
			"*",
		},
		AllowCredentials: true,
	}).Handler(mux)

	fmt.Println("Server is running and listening at port 8080...")
	http.ListenAndServe(":8080", c)
}
