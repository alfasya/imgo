package router

import (
	"github.com/alfasya/imgo/handler"
	"github.com/gin-gonic/gin"
)

func File(r *gin.Engine) {
	file := r.Group("/auth")
	file.POST("/upload", handler.Upload)
	file.POST("/edit", handler.Edit)
	file.DELETE("/delete", handler.Delete)
}
