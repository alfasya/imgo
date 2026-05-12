package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/alfasya/imgo/utils"
)

type Image struct {
	Name      string
	Size      int64
	Date      string
	Thumbnail string
}

type GalleryRes struct {
	Message   string
	Owner     utils.Owner
	ImageList []Image
}

func Gallery(w http.ResponseWriter, r *http.Request) {
	//parse token from request header
	tokenString := r.Header.Get("Authorization")

	//handling empty authorization header
	if tokenString == "" {
		http.Error(w, "missing token", http.StatusUnauthorized)
		return
	}

	stringParts := strings.Split(tokenString, " ")

	if len(stringParts) == 0 || strings.ToLower(stringParts[0]) != "bearer" {
		http.Error(w, "missing bearer token", http.StatusUnauthorized)
		return
	}

	tokenString = stringParts[1]

	owner, err := utils.VerifyToken(tokenString)
	if err != nil {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(GalleryRes{
		Message:   fmt.Sprintf("Welcome %s", owner.Username),
		Owner:     owner,
		ImageList: []Image{},
	})
}
