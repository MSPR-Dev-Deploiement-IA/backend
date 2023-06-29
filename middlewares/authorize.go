package middlewares

import (
	"backend/security"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (m Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the access token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			return
		}

		// Validate the access token
		userID, err := security.ValidateAccessToken(splitToken[1])
		if err != nil {
			log.Printf("Error validating access token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		// Set the user ID in the context
		c.Set("userID", userID)
		c.Next()
	}
}
