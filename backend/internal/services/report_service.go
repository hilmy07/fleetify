package services

import (
	"fleetify/internal/config"
	"fleetify/internal/dto"
	"fleetify/internal/models"
	"fleetify/internal/repositories"

	"gorm.io/gorm"
)

type ReportService struct {
	Repo *repositories.ReportRepository
}

func NewReportService(repo *repositories.ReportRepository) *ReportService {
	return &ReportService{Repo: repo}
}

func (s *ReportService) CreateReport(req dto.CreateReportRequest, userID uint) error {

	return config.DB.Transaction(func(tx *gorm.DB) error {

		report := models.MaintenanceReport{
			VehicleID: req.VehicleID,
			CreatedBy: userID,
			Odometer: req.Odometer,
			Complaint: req.Complaint,
			ProofPhoto: req.InitialPhoto,
			Status: "PENDING_APPROVAL",
		}

		if err := s.Repo.CreateReport(tx, &report); err != nil {
			return err
		}

		for _, item := range req.Items {

			masterItem, err := s.Repo.FindMasterItemByID(item.ItemID)
			if err != nil {
				return err
			}

			reportItem := models.ReportItem{
				ReportId: report.Id,
				ItemId: item.ItemID,
				Quantity: item.Quantity,
				PriceSnapshot: masterItem.Price,
			}

			if err := s.Repo.CreateReportItem(tx, &reportItem); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *ReportService) ApproveReport(id uint) error {

	report, err := s.Repo.FindReportByID(id)
	if err != nil {
		return err
	}

	report.Status = "APPROVED"

	return s.Repo.Save(report)
}

func (s *ReportService) CompleteReport(id uint, proofPhoto string) error {

	report, err := s.Repo.FindReportByID(id)
	if err != nil {
		return err
	}

	report.Status = "COMPLETED"
	report.ProofPhoto = proofPhoto

	return s.Repo.Save(report)
}

func (s *ReportService) GetAll() ([]models.MaintenanceReport, error) {
	return s.Repo.GetAll()
}