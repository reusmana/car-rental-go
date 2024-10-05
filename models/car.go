package models

type Car struct {
	ID           uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Brand        string  `json:"brand"`
	Model        string  `json:"model"`
	DailyRent    float64 `json:"daily_rent"`
	Availability bool    `json:"availability"`
}
