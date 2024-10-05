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

	membershipID1 := int64(1)
	membershipID2 := int64(2)
	membershipID3 := int64(3)

	customers := []models.Customer{
		{Name: "Wawan Hermawan", NIK: "3372093912739", Phone: "081237123682", MembershipID: &membershipID1},
		{Name: "Philip Walker", NIK: "3372093912785", Phone: "081237123683", MembershipID: &membershipID3},
		{Name: "Hugo Fleming", NIK: "3372093912800", Phone: "081237123684", MembershipID: nil},
		{Name: "Maximillian Mendez", NIK: "3372093912848", Phone: "081237123685", MembershipID: &membershipID2},
		{Name: "Felix Dixon", NIK: "3372093912851", Phone: "081237123686", MembershipID: &membershipID1},
		{Name: "Nicholas Riddle", NIK: "3372093912929", Phone: "081237123687", MembershipID: nil},
		{Name: "Stephen Wheeler", NIK: "3372093912976", Phone: "081237123688", MembershipID: &membershipID1},
		{Name: "Roy Brennan", NIK: "3372093913022", Phone: "081237123689", MembershipID: nil},
		{Name: "Eliza Le", NIK: "3372093913106", Phone: "081237123690", MembershipID: nil},
		{Name: "Jesse Taylor", NIK: "3372093913126", Phone: "081237123691", MembershipID: &membershipID3},
		{Name: "Damien Kaufman", NIK: "3372093913202", Phone: "081237123692", MembershipID: &membershipID1},
		{Name: "Ayesha Richardson", NIK: "3372093913257", Phone: "081237123693", MembershipID: &membershipID2},
		{Name: "Margaret Stokes", NIK: "3372093913262", Phone: "081237123694", MembershipID: nil},
		{Name: "Sara Livingston", NIK: "3372093913268", Phone: "081237123695", MembershipID: nil},
		{Name: "Callie Townsend", NIK: "3372093913281", Phone: "081237123696", MembershipID: nil},
		{Name: "Lilly Fischer", NIK: "3372093913325", Phone: "081237123697", MembershipID: &membershipID3},
		{Name: "Theresa Barton", NIK: "3372093913335", Phone: "081237123698", MembershipID: &membershipID1},
		{Name: "Mia Curtis", NIK: "3372093913343", Phone: "081237123699", MembershipID: nil},
		{Name: "Flora Barlow", NIK: "3372093913400", Phone: "081237123700", MembershipID: &membershipID2},
		{Name: "Vanessa Patton", NIK: "3372093913434", Phone: "081237123701", MembershipID: &membershipID2},
	}
	for _, customer := range customers {
		db.FirstOrCreate(&customer, models.Customer{Name: customer.Name, NIK: customer.NIK, Phone: customer.Phone, MembershipID: customer.MembershipID})
	}

	log.Println("Seeding completed successfully!")
}
