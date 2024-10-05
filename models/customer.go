package models

type Customer struct {
	ID           uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name"`
	NIK          string `json:"nik"`
	Phone        string `json:"phone"`
	MembershipID *int64 `json:"membership"`
}
