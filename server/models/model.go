package models

import "github.com/alfasya/imgo/utils"

type Image struct {
	Name string
	Size int64
	Path string
}

type GalleryRes struct {
	Message   string
	Owner     utils.Owner
	ImageList []Image
}
