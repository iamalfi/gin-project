package routes

import (
	"gin-project/controller/user"
	"gin-project/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {

	userRoutes := r.Group("/user")

	// r.GET("/user", middleware.AuthMiddleware(middleware.User), func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Welcome, user!",
	// 	})
	// })
	userRoutes.POST("/create", user.CreateUser)
	userRoutes.POST("/login", user.LoginUser)
	userRoutes.GET("/all", user.GetUsers)
	userRoutes.GET("/token", middleware.AuthMiddleware(), user.GetUserByToken)
	userRoutes.GET("details/:id", user.GetUserByID)
	userRoutes.PATCH("update/:id", user.UpdateUser)
	userRoutes.DELETE("delete/:id", user.DeleteUser)

}
