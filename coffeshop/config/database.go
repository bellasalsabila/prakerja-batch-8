package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDatabase() {
	// Load variabel lingkungan dari file .env
	errdb := godotenv.Load(".env")
	if errdb != nil {
		panic("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error
	DB, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Migrate database models here
	// DB.AutoMigrate(&YourModel{})
}

func CloseDatabase() {
	DB.Close()
}
