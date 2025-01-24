package user

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
func CreateUser(c *gin.Context) {
	user := model.User{}
	c.Bind(&user)
	if err := database.DB.Table("users").Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.Error(helper.New(http.StatusConflict, "User Already Exist", err))
		return
	}

	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		c.Error(helper.New(http.StatusBadRequest, "Password didn't hashed", err))
		return
	}
	user.Password = hashedPassword
	if err := database.DB.Table("users").Create(&user).Error; err != nil {
		c.Error(helper.New(http.StatusBadRequest, "Failed to create user", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}
