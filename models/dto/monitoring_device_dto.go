package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// MonitoringDeviceCreateDTO is used for creating a new monitoring device
type MonitoringDeviceCreateDTO struct {
	Status string `json:"status"`
}

// MonitoringDeviceUpdateDTO is used for updating an existing monitoring device
type MonitoringDeviceUpdateDTO struct {
	Status     string     `json:"status"`
	PatientID  *uuid.UUID `json:"patient_id"`
	LinkedByID *uuid.UUID `json:"linked_by_id"`
}

// MonitoringDeviceDTO is used for retrieving a monitoring device
type MonitoringDeviceDTO struct {
	DeviceID string               `json:"device_id"`
	Status   string               `json:"status"`
	Patient  *PatientForDeviceDTO `json:"patient"`
	LinkedBy *DoctorDTO           `json:"linked_by"`
}

type MonitoringDeviceFilter struct {
	DNI string `json:"dni"`
}

// MapMonitoringDeviceToDTO maps a MonitoringDevice model to a MonitoringDeviceDTO
func MapMonitoringDeviceToDTO(device *models.MonitoringDevice) *MonitoringDeviceDTO {
	if device.Patient == nil {
		device.Patient = &models.Patient{}
	}

	return &MonitoringDeviceDTO{
		DeviceID: device.DeviceID,
		Status:   device.Status,
		Patient:  MapPatientToPatientForDeviceDTO(device.Patient),
		LinkedBy: MapDoctorToDTO(device.LinkedBy),
	}
}

// MapMonitoringDevicesToDTOs maps a list of MonitoringDevice models to a list of MonitoringDeviceDTOs
func MapMonitoringDevicesToDTOs(devices []*models.MonitoringDevice) []*MonitoringDeviceDTO {
	var deviceDTOs = make([]*MonitoringDeviceDTO, 0)
	for _, device := range devices {
		deviceDTOs = append(deviceDTOs, MapMonitoringDeviceToDTO(device))
	}
	return deviceDTOs
}
