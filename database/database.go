package database

import (
	"Osheet-api/v1/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv" // Use godotenv to get .env
	"gorm.io/driver/postgres"  // Use postgresql driver in gorm
	"gorm.io/gorm"             // Use gorm for orm tool
)

var db *gorm.DB

func init() {

	// Load .env configs
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error: .env File Not Found")
	}

	// Use Getenv to read .env
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	port := os.Getenv("DB_PORT")

	// db connection
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei", dbHost, username, password, dbName, port)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database Connection Failed:", err)
	}

	channel := new(models.Channel)

	db.AutoMigrate(channel)
}

func GetDB() *gorm.DB {
	return db
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}

	sqlDB.Close()
}
