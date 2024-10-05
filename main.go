package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/reusmana/car-rental-go/config"
	routes "github.com/reusmana/car-rental-go/router"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	println("Hello, World!")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env := os.Getenv("DATABASE_URL")
	println(env)

	// ConnectDB()

	e := echo.New()

	// Connect to Database
	config.ConnectDB()

	// Setup routes
	routes.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

func ConnectDB() {
	dsn := os.Getenv("DATABASE_URL") // Example: "host=localhost user=postgres password=pass dbname=car_rental port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to database")
		return
	}

	// Auto migrate the customer model
	// db.AutoMigrate(&models.Customer{})
	DB = db
	fmt.Println("Connected to the database successfully!")
}
