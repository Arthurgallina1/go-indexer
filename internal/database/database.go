package database

import (
	"fmt"
	"log"
	"os"

	"go-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB is a global variable to hold the database connection
var DB *gorm.DB

func ConnectDB() {
	// build the database connection string from environment variables
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

	// connec to db
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database. \n", err)
    }

    log.Println("Connected Successfully to Database")
    
	// AutoMigrate will create/update tables based on models
    db.AutoMigrate(&models.Post{})

	// Save the database connection to our global variable
    DB = db
}