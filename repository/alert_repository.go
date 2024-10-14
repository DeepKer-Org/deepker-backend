package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

// AlertRepository includes specific methods for the Alert entity and embeds BaseRepository
type AlertRepository interface {
	BaseRepository[models.Alert]
	GetAttendedAlerts(offset int, limit int) ([]*models.Alert, error)
	GetUnattendedAlerts(offset int, limit int) ([]*models.Alert, error)
	CountAlertsByStatus(status string, count *int64) error
}

// alertRepository struct embeds baseRepository for common CRUD operations
type alertRepository struct {
	BaseRepository[models.Alert]
	db *gorm.DB
}

// NewAlertRepository creates a new instance of AlertRepository
func NewAlertRepository(db *gorm.DB) AlertRepository {
	baseRepo := NewBaseRepository[models.Alert](db)
	return &alertRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetByID retrieves an alert by its ID.
func (r *alertRepository) GetByID(id interface{}, primaryKey string) (*models.Alert, error) {
	var alert models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Preload("ComputerDiagnostics").
		Where(primaryKey+" = ?", id).First(&alert).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &alert, nil
}

// GetAll retrieves all alerts.
func (r *alertRepository) GetAll() ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Preload("ComputerDiagnostics").
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetAttendedAlerts retrieves attended alerts with pagination (attended_timestamp is not null).
func (r *alertRepository) GetAttendedAlerts(offset int, limit int) ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Preload("ComputerDiagnostics").
		Where("attended_timestamp IS NOT NULL").
		Offset(offset).
		Limit(limit).
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetUnattendedAlerts retrieves unattended alerts with pagination (attended_timestamp is null).
func (r *alertRepository) GetUnattendedAlerts(offset int, limit int) ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Preload("ComputerDiagnostics").
		Where("attended_timestamp IS NULL").
		Offset(offset).
		Limit(limit).
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *alertRepository) CountAlertsByStatus(status string, count *int64) error {
	var condition string
	if status == "attended" {
		condition = "attended_timestamp IS NOT NULL"
	} else {
		condition = "attended_timestamp IS NULL"
	}
	return r.db.Model(&models.Alert{}).Where(condition).Count(count).Error
}
