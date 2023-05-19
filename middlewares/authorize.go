package middlewares

import (
	"backend/security"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (m Middleware) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the access token from the cookie
		token, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access token cookie is required"})
			return
		}

		// Validate the access token
		userID, err := security.ValidateAccessToken(token)
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
