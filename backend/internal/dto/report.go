package dto

type CreateReportItemRequest struct {
	ItemID   uint `json:"item_id"`
	Quantity uint `json:"quantity"`
}

type CreateReportRequest struct {
	VehicleID    uint                      `json:"vehicle_id"`
	Odometer     uint                      `json:"odometer"`
	Complaint    string                    `json:"complaint"`
	InitialPhoto string                    `json:"initial_photo"`
	Items        []CreateReportItemRequest `json:"items"`
}