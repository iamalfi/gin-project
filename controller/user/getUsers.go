package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []model.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.Error(helper.New(http.StatusInternalServerError, "Failed to fetch users", err))

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Users fetched successfully",
		"user":    users,
	})
}
