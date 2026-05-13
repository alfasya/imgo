package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
	"github.com/alfasya/imgo/models"
	"github.com/alfasya/imgo/utils"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	owner, ok := r.Context().Value("owner").(utils.Owner)
	if !ok {
		http.Error(w, "invalid type", http.StatusBadRequest)
		return
	}

	//query: retrieve images from database
	images, err := db.GetImages(owner.UserId)
	if err != nil {
		http.Error(w, "database error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.GalleryRes{
		Message:   fmt.Sprintf("Welcome %s", owner.Username),
		Owner:     owner,
		ImageList: images,
	})
}
