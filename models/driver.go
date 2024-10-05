package models

type Driver struct {
	ID        uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string  `json:"name"`      // Cost per day for the driver
	Incentive float64 `json:"incentive"` // Incentive calculated for the driver
}
