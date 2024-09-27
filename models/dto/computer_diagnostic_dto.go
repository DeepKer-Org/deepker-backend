package dto

import "biometric-data-backend/models"

// ComputerDiagnosticCreateDTO is used for creating a new computer diagnosis
type ComputerDiagnosticCreateDTO struct {
	AlertID    string  `json:"alert_id"`
	Diagnosis  string  `json:"diagnosis"`
	Percentage float64 `json:"percentage"`
}

// ComputerDiagnosticUpdateDTO is used for updating an existing computer diagnosis
type ComputerDiagnosticUpdateDTO struct {
	AlertID    string  `json:"alert_id"`
	Diagnosis  string  `json:"diagnosis"`
	Percentage float64 `json:"percentage"`
}

// ComputerDiagnosticDTO is used for retrieving a computer diagnosis
type ComputerDiagnosticDTO struct {
	DiagnosisID uint    `json:"diagnosis_id"`
	AlertID     string  `json:"alert_id"`
	Diagnosis   string  `json:"diagnosis"`
	Percentage  float64 `json:"percentage"`
}

// MapComputerDiagnosticToDTO maps a ComputerDiagnostic model to a ComputerDiagnosisDTO
func MapComputerDiagnosticToDTO(diagnosis *models.ComputerDiagnostic) *ComputerDiagnosticDTO {
	return &ComputerDiagnosticDTO{
		DiagnosisID: diagnosis.DiagnosisID,
		AlertID:     diagnosis.AlertID,
		Diagnosis:   diagnosis.Diagnosis,
		Percentage:  diagnosis.Percentage,
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
