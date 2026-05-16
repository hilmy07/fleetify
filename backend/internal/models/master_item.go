package models

type MasterItem struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	ItemName string  `gorm:"not null" json:"item_name"`
	Type     string  `gorm:"not null" json:"type"`
	Price    float64 `gorm:"not null" json:"price"`
}