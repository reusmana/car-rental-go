package config

import (
	"fmt"
	"log"
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
	db.AutoMigrate(&models.Car{}, &models.Customer{}, &models.Booking{}, &models.Driver{}, &models.Membership{})
	DB = db
	seed(DB)
	fmt.Println("Connected to the database successfully!")
}

func seed(db *gorm.DB) {
	memberships := []models.Membership{
		{Name: "gold"},
		{Name: "silver"},
		{Name: "bronze"},
	}
	for _, membership := range memberships {
		db.FirstOrCreate(&membership, models.Membership{Name: membership.Name})
	}

	log.Println("Seeding completed successfully!")
}
