package main

import (
	"gin-project/database"
	"gin-project/middleware"
	"gin-project/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.ConnectDB()

	// Set Gin mode to release
	// gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(middleware.ErrorHandler)

	api := r.Group("/api")
	routes.Routes(api)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
