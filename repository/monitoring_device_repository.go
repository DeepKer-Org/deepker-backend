package repository

import (
	"biometric-data-backend/models"
	"biometric-data-backend/models/dto"
	"gorm.io/gorm"
)

type MonitoringDeviceRepository interface {
	CreateMonitoringDevice(monitoringDevice *models.MonitoringDevice) error
	GetMonitoringDeviceByID(id string) (*models.MonitoringDevice, error)
	GetAllMonitoringDevices(offset int, limit int, filters dto.MonitoringDeviceFilter) ([]*models.MonitoringDevice, error)
	GetDevicesByStatus(status string) ([]*models.MonitoringDevice, error)
	CountAllMonitoringDevices(filters dto.MonitoringDeviceFilter) (int64, error)
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
	var device models.MonitoringDevice
	// Query by device_id
	if err := r.db.Where("device_id = ?", id).First(&device).Error; err != nil {
		return nil, err
	}
	return &device, nil
}

func applyMonitoringDeviceFilters(query *gorm.DB, filters dto.MonitoringDeviceFilter) *gorm.DB {
	// Join the patients table to filter by patient DNI using partial match (LIKE)
	if filters.DNI != "" {
		query = query.Joins("JOIN patients ON patients.patient_id = monitoring_devices.patient_id").
			Where("patients.dni LIKE ?", "%"+filters.DNI+"%")
	}
	return query
}

func (r *monitoringDeviceRepository) GetAllMonitoringDevices(offset int, limit int, filters dto.MonitoringDeviceFilter) ([]*models.MonitoringDevice, error) {
	var monitoringDevices []*models.MonitoringDevice

	query := r.db.
		Preload("Patient").
		Preload("LinkedBy")
	query = applyMonitoringDeviceFilters(query, filters) // Apply filters

	// Apply pagination
	if err := query.Offset(offset).Limit(limit).Find(&monitoringDevices).Error; err != nil {
		return nil, err
	}

	return monitoringDevices, nil
}

func (r *monitoringDeviceRepository) GetDevicesByStatus(status string) ([]*models.MonitoringDevice, error) {
	var devices []*models.MonitoringDevice
	if err := r.db.
		Preload("Patient").
		Preload("LinkedBy").
		Where("status = ?", status). // Filter by status
		Find(&devices).Error; err != nil {
		return nil, err
	}
	return devices, nil
}

func (r *monitoringDeviceRepository) CountAllMonitoringDevices(filters dto.MonitoringDeviceFilter) (int64, error) {
	var totalCount int64

	query := r.db.Model(&models.MonitoringDevice{})
	query = applyMonitoringDeviceFilters(query, filters) // Apply filters

	// Count the total number of devices with the filters applied
	if err := query.Count(&totalCount).Error; err != nil {
		return 0, err
	}

	return totalCount, nil
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
	if err := r.db.Delete(&models.MonitoringDevice{}).Error; err != nil {
		return err
	}
	return nil
}
