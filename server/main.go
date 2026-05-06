package main

import (
	"github.com/alfasya/imgo/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	router.Auth(r)
	router.User(r)

	r.Run("localhost:8080")
}
