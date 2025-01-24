package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserByToken(c *gin.Context) {

	user := model.User{}
	userID := c.MustGet("id").(string)

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	if err := database.DB.Table("users").Where("id = ?", parsedUserID).First(&user).Error; err != nil {
		c.Error(helper.New(http.StatusNotFound, "User Not found", err))
		return
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}
