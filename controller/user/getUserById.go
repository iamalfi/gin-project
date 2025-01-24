package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUserByID(c *gin.Context) {
	id := c.Param("id")

	user := model.User{}
	err := database.DB.Table("users").Where("id = ?", id).First(&user).Error

	if err != nil || user.ID == uuid.Nil {
		c.Error(helper.New(http.StatusNotFound, "User Not Found", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User fetched successfully",
		"user":    user,
	})
}
