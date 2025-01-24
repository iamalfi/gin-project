package routes

import (
	"gin-project/controller/product"
	"gin-project/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup) {
	productRoutes := r.Group("/product")

	productRoutes.POST("/create", middleware.AuthMiddleware(middleware.Client), product.CreateProduct)

}
