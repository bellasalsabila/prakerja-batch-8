package config

import (
	"coffeshop/model"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
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
	initMigration()
}

func CloseDatabase() {
	DB.Close()
}

func initMigration() {

	DB.AutoMigrate(&model.Data{}, &model.User{})
}
