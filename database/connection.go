package database

import (
	"fmt"
	"gin-project/model"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	// Load database URL from environment variables
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		panic("DATABASE_URL environment variable is not set")
	}
	fmt.Println("DATABASE_URL environment variable is set...")

	// Establish DB connection
	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Discard.LogMode(logger.Error),
	})
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %v", err))
	}
	fmt.Println("Connected to DB")

	// Perform AutoMigrate for User model
	migrationErr := DB.AutoMigrate(
		&model.User{},
		&model.Product{},
	)
	if migrationErr != nil {
		panic(fmt.Sprintf("Error migrating database: %v", migrationErr))
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to get underlying sql.DB: %v", err))
	}

	sqlDB.SetMaxOpenConns(100)  // Maximum number of open connections
	sqlDB.SetMaxIdleConns(10)   // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(0) // Maximum amount of time a connection can be reused (0 means no limit)

	fmt.Println("DB Connection pool configured for 100 connections")
	fmt.Println("DB Migration complete")
}
