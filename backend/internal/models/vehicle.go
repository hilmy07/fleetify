package models

type Vehicle struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	LicensePlate string `gorm:"not null" json:"license_plate"`
	Model        string `gorm:"not null" json:"model"`
}
