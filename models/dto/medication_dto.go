package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)
import "time"

// MedicationCreateDTO is used for creating a new medication
type MedicationCreateDTO struct {
	PatientID   uuid.UUID  `json:"patient_id"`
	Name        string     `json:"name"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Dosage      string     `json:"dosage"`
	Periodicity string     `json:"periodicity"`
}

// MedicationUpdateDTO is used for updating an existing medication
type MedicationUpdateDTO struct {
	PatientID   uuid.UUID  `json:"patient_id"`
	Name        string     `json:"name"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Dosage      string     `json:"dosage"`
	Periodicity string     `json:"periodicity"`
}

// ShortMedicationDTO is used for retrieving a medication
type ShortMedicationDTO struct {
	Name        string     `json:"name"`
	StartDate   *time.Time `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	Dosage      string     `json:"dosage"`
	Periodicity string     `json:"periodicity"`
}

// MedicationDTO is used for retrieving a medication
type MedicationDTO struct {
	MedicationID uuid.UUID `json:"medication_id"`
	PatientID    uuid.UUID `json:"patient_id"`
	ShortMedicationDTO
}

// MapCommonMedicationFields maps common fields of a Medication model to a ShortMedicationDTO
func MapCommonMedicationFields(medication *models.Medication) ShortMedicationDTO {
	return ShortMedicationDTO{
		Name:        medication.Name,
		StartDate:   medication.StartDate,
		EndDate:     medication.EndDate,
		Dosage:      medication.Dosage,
		Periodicity: medication.Periodicity,
	}
}

// MapShortMedicationToDTO maps a Medication model to a ShortMedicationDTO
func MapShortMedicationToDTO(medication *models.Medication) *ShortMedicationDTO {
	commonFields := MapCommonMedicationFields(medication)
	return &commonFields
}

// MapShortMedicationsToDTOs maps a list of Medication models to a list of ShortMedicationDTOs
func MapShortMedicationsToDTOs(medications []*models.Medication) []*ShortMedicationDTO {
	var medicationDTOs []*ShortMedicationDTO
	for _, medication := range medications {
		medicationDTOs = append(medicationDTOs, MapShortMedicationToDTO(medication))
	}
	return medicationDTOs
}

// MapMedicationToDTO maps a Medication model to a MedicationDTO
func MapMedicationToDTO(medication *models.Medication) *MedicationDTO {
	commonFields := MapCommonMedicationFields(medication)
	return &MedicationDTO{
		MedicationID:       medication.MedicationID,
		PatientID:          medication.PatientID,
		ShortMedicationDTO: commonFields,
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

func MapMedicationsToMedicationsDetails(medications []*models.Medication) []string {
	var medicationDetails []string
	for _, medication := range medications {
		medicationDetails = append(medicationDetails, medication.Name+" - "+medication.Dosage+" - "+medication.Periodicity)
	}
	return medicationDetails
}
