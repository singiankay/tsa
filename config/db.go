package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/singiankay/tsa/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	hostname := os.Getenv("POSTGRES_HOST")
	dbname := os.Getenv("POSTGRES_DB")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")

	dsn := "host="+ hostname +" user="+username+" password="+password+" dbname="+dbname+" port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.Contact{}, &models.PhoneNumber{})
	DB = db
}