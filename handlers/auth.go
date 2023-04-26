package handlers

import (
	"backend/models"
	"backend/security"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strings"
)

func (h Handler) Register(c *gin.Context) {
	// Bind the JSON body to a struct to get the data
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the user to the database, hashing the password, and return the user
	if err := user.Save(h.db); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove the password from the response
	user.Password = ""

	// Return the user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (h Handler) Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userFromDB models.User
	userFromDB.Email = user.Email
	if err := userFromDB.GetUser(h.db); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Compare the password hash in the database with the password entered by the user
	if err := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(user.Password)); err != nil {
		log.Printf("Error comparing passwords: %v", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate the access and refresh tokens
	accessToken, err := security.GenerateAccessToken(uint(userFromDB.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := security.GenerateRefreshToken(uint(userFromDB.ID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	// Add the tokens to the response headers
	c.Header("Access-Token", accessToken)
	c.Header("Refresh-Token", refreshToken)

	// Remove the password from the response
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"user":          user,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (h Handler) Refresh(c *gin.Context) {
	refreshToken := c.Query("refresh_token")

	if refreshToken == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	userID, err := security.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	// Generate a new access token
	accessToken, err := security.GenerateAccessToken(userID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}

func (h Handler) Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			return
		}

		userID, err := security.ValidateAccessToken(token)
		if err != nil {
			log.Printf("Error validating access token: %v", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			return
		}

		c.Set("userID", userID)
		c.Next()
	}
}
