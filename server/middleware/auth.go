package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Header struct {
	Authorization string
}

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header Header

		err := c.ShouldBindHeader(&header)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid header"})
			return
		}

		if header.Authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		parts := strings.Split(header.Authorization, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		key := parts[1]

		fmt.Println(key)
		//vVERIFY THE KEY  TO THE ACCOUNT AUTHORIZATION WITH JWT

		c.Next()
	}
}
