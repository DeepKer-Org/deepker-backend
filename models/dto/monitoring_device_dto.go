package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// MonitoringDeviceCreateDTO is used for creating a new monitoring device
type MonitoringDeviceCreateDTO struct {
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	PatientID uuid.UUID `json:"patient_id"`
	Sensors   []string  `json:"sensors"`
}

// MonitoringDeviceUpdateDTO is used for updating an existing monitoring device
type MonitoringDeviceUpdateDTO struct {
	Status    string    `json:"status"`
	PatientID uuid.UUID `json:"patient_id"`
	Sensors   []string  `json:"sensors"`
}

// MonitoringDeviceDTO is used for retrieving a monitoring device
type MonitoringDeviceDTO struct {
	DeviceID  uuid.UUID `json:"device_id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	PatientID uuid.UUID `json:"patient_id"`
	Sensors   []string  `json:"sensors"`
}

// MapMonitoringDeviceToDTO maps a MonitoringDevice model to a MonitoringDeviceDTO
func MapMonitoringDeviceToDTO(device *models.MonitoringDevice) *MonitoringDeviceDTO {
	return &MonitoringDeviceDTO{
		DeviceID:  device.DeviceID,
		Type:      device.Type,
		Status:    device.Status,
		PatientID: device.PatientID,
		Sensors:   device.Sensors,
	}
}

// MapMonitoringDevicesToDTOs maps a list of MonitoringDevice models to a list of MonitoringDeviceDTOs
func MapMonitoringDevicesToDTOs(devices []*models.MonitoringDevice) []*MonitoringDeviceDTO {
	var deviceDTOs []*MonitoringDeviceDTO
	for _, device := range devices {
		deviceDTOs = append(deviceDTOs, MapMonitoringDeviceToDTO(device))
	}
	return deviceDTOs
}
