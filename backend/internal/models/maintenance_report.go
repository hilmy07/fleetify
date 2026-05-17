package models

type MaintenanceReport struct {
	Id          uint   `gorm:"primaryKey" json:"id"`
	VehicleID   uint   `gorm:"not null" json:"vehicle_id"`
	CreatedBy   uint   `gorm:"not null" json:"created_by"`
	Odometer    uint   `gorm:"not null" json:"odometer"`
	Complaint   string `gorm:"not null" json:"complaint"`
	Status      string `gorm:"not null" json:"status"`
	InitialPost string `gorm:"not null" json:"initial_post"`
	ProofPhoto  string `gorm:"not null" json:"proof_photo"`
	CreatedAt   string `gorm:"not null" json:"created_at"`

	Vehicle Vehicle      `gorm:"foreignKey:VehicleID" json:"vehicle"`
	Items   []ReportItem `gorm:"foreignKey:ReportId" json:"items"`
}