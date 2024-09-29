package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"time"
)

// AlertCreateDTO is used for creating a new alert
type AlertCreateDTO struct {
	AlertStatus    string    `json:"alert_status"`
	Room           string    `json:"room"`
	AlertTimestamp time.Time `json:"alert_timestamp"`
	PatientID      uuid.UUID `json:"patient_id"`
}

// AlertUpdateDTO is used for updating an existing alert
type AlertUpdateDTO struct {
	AlertStatus       string     `json:"alert_status"`
	Room              string     `json:"room"`
	AttendedTimestamp *time.Time `json:"attended_timestamp"`
	PatientID         uuid.UUID  `json:"patient_id"`
}

// AlertDTO is used for retrieving an alert along with related entities
type AlertDTO struct {
	AlertID             uuid.UUID                `json:"alert_id"`
	AlertTimestamp      time.Time                `json:"alert_timestamp"`
	Room                string                   `json:"room"`
	AlertStatus         string                   `json:"alert_status"`
	AttendedBy          string                   `json:"attended_by,omitempty"`
	AttendedTimestamp   *time.Time               `json:"attended_timestamp,omitempty"`
	BiometricData       *BiometricDataDTO        `json:"biometric_data"`
	ComputerDiagnostics []*ComputerDiagnosticDTO `json:"computer_diagnoses"`
	Patient             *PatientForAlertDTO      `json:"patient"`
}

// MapAlertToDTO maps an Alert model to an AlertDTO
func MapAlertToDTO(alert *models.Alert) *AlertDTO {
	var doctorName string
	if alert.AttendedBy != nil {
		doctorName = alert.AttendedBy.Name
	}
	if alert.BiometricData == nil {
		alert.BiometricData = &models.BiometricData{}
	}
	if alert.Patient == nil {
		alert.Patient = &models.Patient{}
	}
	return &AlertDTO{
		AlertID:             alert.AlertID,
		AlertStatus:         alert.AlertStatus,
		Room:                alert.Room,
		AlertTimestamp:      alert.AlertTimestamp,
		AttendedTimestamp:   alert.AttendedTimestamp,
		AttendedBy:          doctorName,
		BiometricData:       MapBiometricDataToDTO(alert.BiometricData),
		ComputerDiagnostics: MapComputerDiagnosticsToDTOs(alert.ComputerDiagnostics),
		Patient:             MapPatientToPatientForAlertDTO(alert.Patient),
	}
}

// MapAlertsToDTOs maps a list of Alert models to a list of AlertDTOs
func MapAlertsToDTOs(alerts []*models.Alert) []*AlertDTO {
	var alertDTOs []*AlertDTO
	for _, alert := range alerts {
		alertDTOs = append(alertDTOs, MapAlertToDTO(alert))
	}
	return alertDTOs
}
