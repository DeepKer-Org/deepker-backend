package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

// AlertRepository includes specific methods for the Alert entity and embeds BaseRepository
type AlertRepository interface {
	BaseRepository[models.Alert]
	GetAttendedAlerts() ([]*models.Alert, error)
	GetUnattendedAlerts() ([]*models.Alert, error)
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
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetAttendedAlerts retrieves alerts that have been attended (attended_timestamp is not null).
func (r *alertRepository) GetAttendedAlerts() ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Where("attended_timestamp IS NOT NULL").Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetUnattendedAlerts retrieves alerts that have not been attended (attended_timestamp is null).
func (r *alertRepository) GetUnattendedAlerts() ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Doctors").
		Where("attended_timestamp IS NULL").Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}
