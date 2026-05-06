package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"Message": "/upload",
	})
}

func Edit(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"Message": "/edit",
	})
}

func Delete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/delete",
	})
}
