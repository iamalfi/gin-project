package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Table("users").Delete(model.User{}, "id = ?", id).Error; err != nil {
		c.Error(helper.New(http.StatusInternalServerError, "Failed to delete user", err))
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
