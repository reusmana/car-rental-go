package models

type Booking struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	CustomerID uint    `json:"customer_id"`
	CarID      uint    `json:"car_id"`
	StartDate  string  `json:"start_date"`
	EndDate    string  `json:"end_date"`
	DayOfRent  int64   `json:"day_of_rent"`
	TotalCost  float64 `json:"total_cost"`
	Status     bool    `json:"status" gorm:"default:false"`
}
