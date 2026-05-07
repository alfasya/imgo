package handler

import (
	"fmt"
	"net/http"

	"github.com/alfasya/imgo/db"
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
	exist, err := db.DbValidation(user.Username)
	if exist == false {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	//VALIDATING PASSWORD
	valid, err := db.PasswordValidation(user.Username, user.Password)
	if valid == false {
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
		if user.Username == "" || user.Password == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "please input username and password"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
	}

	exist, err := db.DbValidation(user.Username)
	if exist == true {
		c.JSON(http.StatusNotFound, gin.H{"message": "username already exists"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	//hash password
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		fmt.Printf("Error hashing password: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	//db.Register
	if err := db.NewRegistration(user.Username, hash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internala server error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "registered",
		"User":    user.Username,
	})
}
