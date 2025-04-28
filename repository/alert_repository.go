package repository

import (
	"biometric-data-backend/models"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

// AlertRepository includes specific methods for the Alert entity and embeds BaseRepository
type AlertRepository interface {
	BaseRepository[models.Alert]
	GetRecentAlerts(offset int, limit int) ([]*models.Alert, error)
	GetPastAlerts(offset int, limit int) ([]*models.Alert, error)
	CountAlertsByPeriod(period string, count *int64) error
	GetAlertsByTimezone(timezone string) ([]*models.Alert, error)
	Liberate(alert *models.Alert) error
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
		Preload("Patient.Comorbidities").
		Preload("Patient.Medications").
		Preload("Patient.Doctors").
		Preload("Patient.MonitoringDevice").
		Preload("ComputerDiagnostic").
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
		Preload("Patient.Comorbidities").
		Preload("Patient.Medications").
		Preload("Patient.Doctors").
		Preload("Patient.MonitoringDevice").
		Preload("ComputerDiagnostic").
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetRecentAlerts retrieves recent alerts with pagination (alert_timestamp <= 24 hours).
func (r *alertRepository) GetRecentAlerts(offset int, limit int) ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Comorbidities").
		Preload("Patient.Medications").
		Preload("Patient.Doctors").
		Preload("ComputerDiagnostic").
		Where("alert_timestamp >= ?", time.Now().UTC().Add(-24*time.Hour)).
		Order("alert_timestamp DESC").
		Offset(offset).
		Limit(limit).
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetPastAlerts retrieves past alerts with pagination (alert_timestamp > 24 hours).
func (r *alertRepository) GetPastAlerts(offset int, limit int) ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Comorbidities").
		Preload("Patient.Medications").
		Preload("Patient.Doctors").
		Preload("Patient.MonitoringDevice").
		Preload("ComputerDiagnostic").
		Where("alert_timestamp < ?", time.Now().UTC().Add(-24*time.Hour)).
		Order("alert_timestamp DESC").
		Offset(offset).
		Limit(limit).
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *alertRepository) GetAlertsByTimezone(timezone string) ([]*models.Alert, error) {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Printf("Invalid timezone: %v", err)
		return nil, err
	}

	// Calculate the start and end of the day in the specified timezone
	now := time.Now().In(loc)
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, loc)
	endOfDay := startOfDay.Add(24 * time.Hour)

	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Preload("Patient.Comorbidities").
		Preload("Patient.Medications").
		Preload("Patient.Doctors").
		Preload("Patient.MonitoringDevice").
		Preload("ComputerDiagnostic").
		Where("alert_timestamp AT TIME ZONE 'UTC' AT TIME ZONE ? >= ? AND alert_timestamp AT TIME ZONE 'UTC' AT TIME ZONE ? < ?", timezone, startOfDay, timezone, endOfDay).
		Order("alert_timestamp DESC").
		Find(&alerts).Error; err != nil {
		return nil, err
	}

	return alerts, nil
}

func (r *alertRepository) CountAlertsByPeriod(period string, count *int64) error {
	var condition string
	if period == "recent" {
		condition = "alert_timestamp >= ?"
	} else {
		condition = "alert_timestamp < ?"
	}

	cutoffTime := time.Now().UTC().Add(-24 * time.Hour)

	return r.db.Model(&models.Alert{}).Where(condition, cutoffTime).Count(count).Error
}

func (r *alertRepository) Liberate(alert *models.Alert) error {
	err := r.db.Model(&models.Alert{}).Where("alert_id = ?", alert.AlertID).Update("attended_timestamp", nil).Update("attended_by_id", nil).Error
	if err != nil {
		return err
	}

	return nil
}
