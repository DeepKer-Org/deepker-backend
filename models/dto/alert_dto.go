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
	AttendedTimestamp *time.Time `json:"attended_timestamp"`
	AttendedByID      uuid.UUID  `json:"attended_by_id"`
}

// AlertDTO is used for retrieving an alert along with related entities
type AlertDTO struct {
	AlertID             uuid.UUID                `json:"alert_id"`
	AlertTimestamp      time.Time                `json:"alert_timestamp"`
	AttendedBy          *DoctorDTO               `json:"attended_by"`
	AttendedTimestamp   string                   `json:"attended_timestamp"`
	AlertStatus         string                   `json:"alert_status"`
	FinalDiagnosis      string                   `json:"final_diagnosis"`
	BiometricData       *BiometricDataDTO        `json:"biometric_data"`
	ComputerDiagnostics []*ComputerDiagnosticDTO `json:"computer_diagnoses"`
	Patient             *PatientForAlertDTO      `json:"patient"`
}

// MapAlertToDTO maps an Alert model to an AlertDTO
func MapAlertToDTO(alert *models.Alert) *AlertDTO {
	if alert.BiometricData == nil {
		alert.BiometricData = &models.BiometricData{}
	}
	if alert.Patient == nil {
		alert.Patient = &models.Patient{}
	}

	if alert.AttendedBy == nil {
		alert.AttendedBy = &models.Doctor{}
	}

	attendedTimestamp := ""
	alertStatus := "Unattended"
	if alert.AttendedTimestamp != nil {
		attendedTimestamp = alert.AttendedTimestamp.Format(time.RFC3339)
		alertStatus = "Attended"
	}

	return &AlertDTO{
		AlertID:             alert.AlertID,
		AlertTimestamp:      alert.AlertTimestamp,
		AttendedTimestamp:   attendedTimestamp,
		AlertStatus:         alertStatus,
		AttendedBy:          MapDoctorToDTO(alert.AttendedBy),
		FinalDiagnosis:      alert.FinalDiagnosis,
		BiometricData:       MapBiometricDataToDTO(alert.BiometricData),
		ComputerDiagnostics: MapComputerDiagnosticsToDTOs(alert.ComputerDiagnostics),
		Patient:             MapPatientToPatientForAlertDTO(alert.Patient),
	}
}

// MapAlertsToDTOs maps a list of Alert models to a list of AlertDTOs
func MapAlertsToDTOs(alerts []*models.Alert) []*AlertDTO {
	alertDTOs := make([]*AlertDTO, 0)
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
