package config

import (
	"fmt"
	"os"

	"github.com/reusmana/car-rental-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	// Auto migrate the customer model
	db.AutoMigrate(&models.Car{}, &models.Customer{}, &models.Booking{})
	DB = db
	fmt.Println("Connected to the database successfully!")
}
