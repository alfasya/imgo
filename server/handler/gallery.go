package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	owner, ok := r.Context().Value("owner").(utils.Owner)
	if !ok {
		http.Error(w, "invalid type", http.StatusBadRequest)
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
