package seeder

import (
	"fleetify/internal/config"
	"fleetify/internal/models"
)

func Seed() {

	// =========================
	// USERS
	// =========================

	users := []models.User{
		{
			Username: "service_advisor",
			Password: "password123",
			Role:     "SA",
		},
		{
			Username: "approval_manager",
			Password: "password123",
			Role:     "APPROVAL",
		},
	}

	for _, user := range users {
		config.DB.FirstOrCreate(
			&user,
			models.User{
				Username: user.Username,
			},
		)
	}

	// =========================
	// VEHICLES
	// =========================

	vehicles := []models.Vehicle{
		{
			LicensePlate: "B 1234 ABC",
			Model:        "Toyota Avanza",
		},
		{
			LicensePlate: "L 5678 DEF",
			Model:        "Honda Brio",
		},
		{
			LicensePlate: "N 9012 GHI",
			Model:        "Mitsubishi Xpander",
		},
	}

	for _, vehicle := range vehicles {
		config.DB.FirstOrCreate(
			&vehicle,
			models.Vehicle{
				LicensePlate: vehicle.LicensePlate,
			},
		)
	}

	// =========================
	// MASTER ITEMS
	// =========================

	items := []models.MasterItem{
		{
			ItemName: "Engine Oil",
			Type:     "PART",
			Price:    250000,
		},
		{
			ItemName: "Oil Filter",
			Type:     "PART",
			Price:    75000,
		},
		{
			ItemName: "Brake Pad",
			Type:     "PART",
			Price:    350000,
		},
		{
			ItemName: "General Service",
			Type:     "SERVICE",
			Price:    500000,
		},
		{
			ItemName: "Wheel Alignment",
			Type:     "SERVICE",
			Price:    200000,
		},
	}

	for _, item := range items {
		config.DB.FirstOrCreate(
			&item,
			models.MasterItem{
				ItemName: item.ItemName,
			},
		)
	}

	// =========================
	// MAINTENANCE REPORT
	// =========================

	report := models.MaintenanceReport{
		VehicleID:   1,
		CreatedBy:   1,
		Odometer:    45200,
		Complaint:   "Engine vibration and noisy braking",
		Status:      "PENDING_APPROVAL",
		InitialPost: "vehicle-check.jpg",
		ProofPhoto:  "",
		CreatedAt:   "2026-05-17",
	}

	config.DB.FirstOrCreate(
		&report,
		models.MaintenanceReport{
			VehicleID: report.VehicleID,
			Complaint: report.Complaint,
		},
	)

	// =========================
	// REPORT ITEMS
	// =========================

	reportItems := []models.ReportItem{
		{
			ReportId:      report.Id,
			ItemId:        1,
			Quantity:      1,
			PriceSnapshot: 250000,
		},
		{
			ReportId:      report.Id,
			ItemId:        2,
			Quantity:      1,
			PriceSnapshot: 75000,
		},
		{
			ReportId:      report.Id,
			ItemId:        4,
			Quantity:      1,
			PriceSnapshot: 500000,
		},
	}

	for _, reportItem := range reportItems {
		config.DB.FirstOrCreate(
			&reportItem,
			models.ReportItem{
				ReportId: reportItem.ReportId,
				ItemId:   reportItem.ItemId,
			},
		)
	}
}