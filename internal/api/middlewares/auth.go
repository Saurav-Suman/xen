package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if authorizationHeader != "secrettoken" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "INVALID_API_KEY"})
			return
		} else {
			c.Next()
		}
	}
}
