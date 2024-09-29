package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// ComputerDiagnosticCreateDTO is used for creating a new computer diagnosis
type ComputerDiagnosticCreateDTO struct {
	AlertID    uuid.UUID `json:"alert_id"`
	Diagnosis  string    `json:"diagnosis"`
	Percentage float64   `json:"percentage"`
}

// ComputerDiagnosticUpdateDTO is used for updating an existing computer diagnosis
type ComputerDiagnosticUpdateDTO struct {
	AlertID    uuid.UUID `json:"alert_id"`
	Diagnosis  string    `json:"diagnosis"`
	Percentage float64   `json:"percentage"`
}

// ComputerDiagnosticDTO is used for retrieving a computer diagnosis
type ComputerDiagnosticDTO struct {
	Diagnosis  string  `json:"diagnosis"`
	Percentage float64 `json:"percentage"`
}

// MapComputerDiagnosticToDTO maps a ComputerDiagnostic model to a ComputerDiagnosisDTO
func MapComputerDiagnosticToDTO(diagnosis *models.ComputerDiagnostic) *ComputerDiagnosticDTO {
	return &ComputerDiagnosticDTO{
		Diagnosis:  diagnosis.Diagnosis,
		Percentage: diagnosis.Percentage,
	}
}

// MapComputerDiagnosticsToDTOs maps a list of ComputerDiagnostic models to a list of ComputerDiagnosisDTOs
func MapComputerDiagnosticsToDTOs(diagnoses []*models.ComputerDiagnostic) []*ComputerDiagnosticDTO {
	var diagnosisDTOs []*ComputerDiagnosticDTO
	for _, diagnosis := range diagnoses {
		diagnosisDTOs = append(diagnosisDTOs, MapComputerDiagnosticToDTO(diagnosis))
	}
	return diagnosisDTOs
}
