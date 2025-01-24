package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.Error(helper.New(http.StatusBadRequest, "Invalid request", err))
		return
	}

	var user model.User
	if err := database.DB.Where("email = ?", loginData.Email).First(&user).Error; err != nil {
		c.Error(helper.New(http.StatusNotFound, "Account doesn't exist", err))
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginData.Password))
	if err != nil {
		c.Error(helper.New(http.StatusBadRequest, "Invalid email or password", err))
		return
	}

	token, err := helper.GenerateToken(user.ID.String(), user.Email, string(user.Role))
	if err != nil {
		c.Error(helper.New(http.StatusBadRequest, "Could not generate token", err))

		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
		"user":    user,
	})
}
