package handler

import (
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/utils"
	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var user User

	//PROCESS JSON AND HANDLING ERROR
	//DESERIALIZING(UNSTRUCTURING) JSON INTO USER STRUCT
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	//SANITIZING JSON
	//VALIDATING JSON
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	//VALIDATING USERNAME FROM DATABASE
	exist, err := DbValidation(user.Username)
	if exist == false {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	//VALIDATING PASSWORD
	valid, err := PasswordValidation(user.Username, user.Password)
	if !valid {
		fmt.Printf("Error validating password: %s", err)
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
	}

	//GENERATING TOKEN
	tokenString, err := utils.GenerateToken(user.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
	}

	//RESPONSE WITH STATUS OK
	c.JSON(http.StatusOK, gin.H{
		"Message": "logged in",
		"Data":    tokenString,
	})
}

func Register(c *gin.Context) {
	var user User

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "/register",
	})
}
