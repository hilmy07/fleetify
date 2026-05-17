package repositories

import (
	"fleetify/internal/models"

	"gorm.io/gorm"
)

type ReportRepository struct {
	DB *gorm.DB
}

func NewReportRepository(db *gorm.DB) *ReportRepository {
	return &ReportRepository{DB: db}
}

func (r *ReportRepository) CreateReport(tx *gorm.DB, report *models.MaintenanceReport) error {
	return tx.Create(report).Error
}

func (r *ReportRepository) CreateReportItem(tx *gorm.DB, item *models.ReportItem) error {
	return tx.Create(item).Error
}

func (r *ReportRepository) FindMasterItemByID(id uint) (*models.MasterItem, error) {
	var item models.MasterItem

	err := r.DB.First(&item, id).Error
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *ReportRepository) FindReportByID(id uint) (*models.MaintenanceReport, error) {
	var report models.MaintenanceReport

	err := r.DB.Preload("Items").First(&report, id).Error
	if err != nil {
		return nil, err
	}

	return &report, nil
}

func (r *ReportRepository) Save(report *models.MaintenanceReport) error {
	return r.DB.Save(report).Error
}

func (r *ReportRepository) GetAll() ([]models.MaintenanceReport, error) {
	var reports []models.MaintenanceReport

	err := r.DB.
		Preload("Vehicle").
		Preload("Items").
		Find(&reports).Error

	return reports, err
}