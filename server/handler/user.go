package handler

import (
	"net/http"

	"github.com/alfasya/imgo/mock"
	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {
	mock.NewUser = &mock.User{
		Id:       212000112371921321,
		Username: "killuazoldyck",
		Role:     "user",
	}

	c.JSON(http.StatusOK, gin.H{
		"Message": "Welcome to the profile page",
		"Data":    mock.NewUser,
	})
}

func Gallery(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "/gallery",
	})
}
