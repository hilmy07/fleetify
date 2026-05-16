package models

type ReportItem struct {
	Id            uint    `gorm:"primaryKey" json:"id"`
	ReportId      uint    `gorm:"not null" json:"report_id"`
	ItemId        uint    `gorm:"not null" json:"item_id"`
	Quantity      uint    `gorm:"not null" json:"quantity"`
	PriceSnapshot float64 `gorm:"not null" json:"price_snapshot"`
}

