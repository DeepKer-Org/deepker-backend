package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

type MonitoringDeviceRepository interface {
	CreateMonitoringDevice(monitoringDevice *models.MonitoringDevice) error
	GetMonitoringDeviceByID(id string) (*models.MonitoringDevice, error)
	GetAllMonitoringDevices() ([]*models.MonitoringDevice, error)
	UpdateMonitoringDevice(monitoringDevice *models.MonitoringDevice) error
	DeleteMonitoringDevice(id string) error
}

type monitoringDeviceRepository struct {
	db *gorm.DB
}

func NewMonitoringDeviceRepository(db *gorm.DB) MonitoringDeviceRepository {
	return &monitoringDeviceRepository{db}
}

// CreateMonitoringDevice creates a new monitoringDevice record in the database.
func (r *monitoringDeviceRepository) CreateMonitoringDevice(monitoringDevice *models.MonitoringDevice) error {
	if err := r.db.Create(monitoringDevice).Error; err != nil {
		return err
	}
	return nil
}

// GetMonitoringDeviceByID retrieves a monitoringDevice by their MonitoringDeviceID.
func (r *monitoringDeviceRepository) GetMonitoringDeviceByID(id string) (*models.MonitoringDevice, error) {
	var monitoringDevice models.MonitoringDevice
	if err := r.db.First(&monitoringDevice, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &monitoringDevice, nil
}

// GetAllMonitoringDevices retrieves all monitoringDevices from the database.
func (r *monitoringDeviceRepository) GetAllMonitoringDevices() ([]*models.MonitoringDevice, error) {
	var monitoringDevices []*models.MonitoringDevice
	if err := r.db.Find(&monitoringDevices).Error; err != nil {
		return nil, err
	}
	return monitoringDevices, nil
}

// UpdateMonitoringDevice updates an existing monitoringDevice record in the database.
func (r *monitoringDeviceRepository) UpdateMonitoringDevice(monitoringDevice *models.MonitoringDevice) error {
	if err := r.db.Save(monitoringDevice).Error; err != nil {
		return err
	}
	return nil
}

// DeleteMonitoringDevice deletes a monitoringDevice by their MonitoringDeviceID.
func (r *monitoringDeviceRepository) DeleteMonitoringDevice(id string) error {
	if err := r.db.Delete(&models.MonitoringDevice{}, id).Error; err != nil {
		return err
	}
	return nil
}
