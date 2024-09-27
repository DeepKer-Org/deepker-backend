package dto

import "biometric-data-backend/models"
import "time"

// MedicationCreateDTO is used for creating a new medication
type MedicationCreateDTO struct {
	PatientID   uint       `json:"patient_id"`
	Medication  string     `json:"medication"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Dosage      string     `json:"dosage"`
	Periodicity string     `json:"periodicity"`
}

// MedicationUpdateDTO is used for updating an existing medication
type MedicationUpdateDTO struct {
	PatientID   uint       `json:"patient_id"`
	Medication  string     `json:"medication"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Dosage      string     `json:"dosage"`
	Periodicity string     `json:"periodicity"`
}

// MedicationDTO is used for retrieving a medication
type MedicationDTO struct {
	MedicationID uint       `json:"medication_id"`
	PatientID    uint       `json:"patient_id"`
	Medication   string     `json:"medication"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Dosage       string     `json:"dosage"`
	Periodicity  string     `json:"periodicity"`
}

// MapMedicationToDTO maps a Medication model to a MedicationDTO
func MapMedicationToDTO(medication *models.Medication) *MedicationDTO {
	return &MedicationDTO{
		MedicationID: medication.MedicationID,
		PatientID:    medication.PatientID,
		Medication:   medication.Medication,
		StartDate:    medication.StartDate,
		EndDate:      medication.EndDate,
		Dosage:       medication.Dosage,
		Periodicity:  medication.Periodicity,
	}
}

// MapMedicationsToDTOs maps a list of Medication models to a list of MedicationDTOs
func MapMedicationsToDTOs(medications []*models.Medication) []*MedicationDTO {
	var medicationDTOs []*MedicationDTO
	for _, medication := range medications {
		medicationDTOs = append(medicationDTOs, MapMedicationToDTO(medication))
	}
	return medicationDTOs
}
