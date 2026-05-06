package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var user User

	//PROCESS JSON AND HANDLING ERROR
	//DESERIALIZING(UNSTRUCTURING) JSON INTO USER STRUCT
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	tokenString, err := GenerateToken(user.Username)
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
	c.JSON(http.StatusCreated, gin.H{
		"Message": "/register",
	})
}
