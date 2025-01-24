package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	user := model.User{}
	if err := database.DB.Table("users").Find(&user, "id = ?", id).Error; err != nil {
		c.Error(helper.New(http.StatusNotFound, "User not found", err))
	}

	c.Bind(&user)

	if err := database.DB.Table("users").Save(user).Error; err != nil {
		c.Error(helper.New(http.StatusInternalServerError, "Failed to update user", err))
	}
	user.Password = ""
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}
