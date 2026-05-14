package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

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
		http.Error(w, fmt.Sprintf("database error: %v", err), http.StatusInternalServerError)
		return
	}

	var links []string
	//compare owner UUID from jwt claims and image owner UUID from database
	if len(images) != 0 {
		imagePath := images[0].Path
		fmt.Println(imagePath)

		imageDir := filepath.Dir(imagePath)
		imageUUID := filepath.Base(imageDir)

		if imageUUID != owner.UserUUID {
			http.Error(w, "forbidden", http.StatusForbidden)
			return
		}

		//create link for each image
		PathLink := (filepath.Join("images", imageUUID))
		for i := range images {
			file := filepath.Base(images[i].Path)
			link := filepath.ToSlash(filepath.Join(PathLink, file))
			links = append(links, link)
		}

		fmt.Println(links)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set(
		"Content-Security-Policy",
		"default-src 'self'; connect-src 'self' http://localhost:5500",
	)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(models.GalleryRes{
		Message:   fmt.Sprintf("Welcome %s", owner.Username),
		Owner:     owner,
		ImageList: images,
		Links:     links,
	})
}
