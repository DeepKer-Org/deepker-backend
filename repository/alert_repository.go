package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type AlertRepository interface {
	CreateAlert(alert *models.Alert) error
	GetAlertByID(id string) (*models.Alert, error)
	GetAllAlerts() ([]*models.Alert, error)
	UpdateAlert(alert *models.Alert) error
	DeleteAlert(id string) error
}

type alertRepository struct {
	db *gorm.DB
}

func NewAlertRepository(db *gorm.DB) AlertRepository {
	return &alertRepository{db}
}

// CreateAlert creates a new alert record in the database.
func (r *alertRepository) CreateAlert(alert *models.Alert) error {
	if err := r.db.Create(alert).Error; err != nil {
		return err
	}
	return nil
}

/* old code
// GetAlertByID retrieves an alert by their AlertID.
func (r *alertRepository) GetAlertByID(id string) (*models.Alert, error) {
	var alert models.Alert
	if err := r.db.First(&alert, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &alert, nil
}
*/
// GetAlertByID retrieves an alert by its AlertID.
func (r *alertRepository) GetAlertByID(id string) (*models.Alert, error) {
	var alert models.Alert
	// Explicitly search by the alert_id field
	if err := r.db.Where("alert_id = ?", id).First(&alert).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &alert, nil
}

// GetAllAlerts retrieves all alerts from the database.
func (r *alertRepository) GetAllAlerts() ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

// UpdateAlert updates an existing alert record in the database.
func (r *alertRepository) UpdateAlert(alert *models.Alert) error {
	if err := r.db.Save(alert).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAlert deletes a alert by their AlertID.
func (r *alertRepository) DeleteAlert(id string) error {
	if err := r.db.Delete(&models.Alert{}, id).Error; err != nil {
		return err
	}
	return nil
}
