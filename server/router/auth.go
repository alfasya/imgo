package router

import (
	"github.com/alfasya/imgo/handler"
	"github.com/gin-gonic/gin"
)

func Auth(r *gin.Engine) {
	auth := r.Group("/auth")
	auth.POST("/register", handler.Register)
	auth.POST("/login", handler.Login)
}
