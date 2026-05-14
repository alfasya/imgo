package handler

import (
	"net/http"
	"path/filepath"
)

func FileServer() http.Handler {
	dir, _ := filepath.Abs("C:/Users/alfas/Documents/images")

	return http.StripPrefix("/images/", http.FileServer(http.Dir(dir)))
}
