package models

type Booking struct {
	ID                 uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID         uint    `json:"customer_id"`
	CarID              uint    `json:"car_id"`
	DriverID           *uint   `json:"driver_id"` // Nullable: optional driver
	RentType           string  `json:"rent_type"`
	StartDate          string  `json:"start_date"`
	EndDate            string  `json:"end_date"`
	DayOfRent          int64   `json:"day_of_rent"`
	TotalCost          float64 `json:"total_cost"`
	MembershipDiscount float64 `json:"membership_discount"`
	DriverIncentive    float64 `json:"driver_incentive,omitempty"`
	Status             bool    `json:"status" gorm:"default:false"`
}
