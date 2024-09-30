package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"time"
)

// AlertCreateDTO is used for creating a new alert
type AlertCreateDTO struct {
	BiometricDataID       uuid.UUID   `json:"biometric_data_id"`
	ComputerDiagnosticIDs []uuid.UUID `json:"computer_diagnostic_ids"`
	PatientID             uuid.UUID   `json:"patient_id"`
}

type AlertCreateResponseDTO struct {
	AlertID string `json:"alert_id,omitempty"`
	Message string `json:"message"`
}

// AlertUpdateDTO is used for updating an existing alert
type AlertUpdateDTO struct {
	Room              string     `json:"room"`
	AttendedTimestamp *time.Time `json:"attended_timestamp"`
	AttendedByID      uuid.UUID  `json:"attended_by_id"`
}

// AlertDTO is used for retrieving an alert along with related entities
type AlertDTO struct {
	AlertID             uuid.UUID                `json:"alert_id"`
	AlertTimestamp      time.Time                `json:"alert_timestamp"`
	Room                string                   `json:"room"`
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
	if alerts == nil {
		return []*AlertDTO{}
	}
	var alertDTOs []*AlertDTO
	for _, alert := range alerts {
		alertDTOs = append(alertDTOs, MapAlertToDTO(alert))
	}
	return alertDTOs
}

func MapCreateDTOToAlert(dto *AlertCreateDTO) *models.Alert {
	return &models.Alert{
		BiometricDataID: dto.BiometricDataID,
		PatientID:       dto.PatientID,
	}
}
