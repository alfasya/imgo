package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/login",
	})
}

func Register(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/register",
	})
}
