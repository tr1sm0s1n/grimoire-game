package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authority() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		expectedToken := "token"
		if authorizationHeader != "Bearer "+expectedToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		c.Next()
	}
}
