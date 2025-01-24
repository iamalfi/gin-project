package product

import (
	"gin-project/database"
	"gin-project/helper"
	"gin-project/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	product := model.Product{}

	c.Bind(&product)
	userID := c.MustGet("id").(string)

	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	product.UserID = parsedUserID

	if err := database.DB.Table("products").Create(&product).Error; err != nil {
		c.Error(helper.New(http.StatusInternalServerError, "Unable to create product", err))
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Product created successfully",
		"product": product,
	})

}
