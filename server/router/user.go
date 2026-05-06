package router

import (
	"github.com/alfasya/imgo/handler"
	"github.com/gin-gonic/gin"
)

func User(r *gin.Engine) {
	user := r.Group("/user")
	user.GET("/profile", handler.Profile)
	user.GET("/gallery", handler.Gallery)
}
