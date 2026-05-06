package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/profile",
	})
}

func Gallery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/gallery",
	})
}
