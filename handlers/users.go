package handlers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetCurrentUser(ctx *gin.Context) {
	userID, _ := ctx.Get("userID")

	var user models.User
	h.db.Where("id = ?", userID).First(&user)

	if user.ID == 0 {
		ctx.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
		return
	}

	ctx.JSON(200, gin.H{"user": user})
}
