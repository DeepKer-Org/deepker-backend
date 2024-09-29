package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// PatientCreateDTO is used for the creation of a new patient
type PatientCreateDTO struct {
	DNI            string  `json:"dni"`
	Name           string  `json:"name"`
	Age            int     `json:"age"`
	Weight         float64 `json:"weight"`
	Height         float64 `json:"height"`
	Sex            string  `json:"sex"`
	Location       string  `json:"location"`
	CurrentState   string  `json:"current_state"`
	FinalDiagnosis string  `json:"final_diagnosis"`
}

// PatientUpdateDTO is used for updating an existing patient
type PatientUpdateDTO struct {
	DNI            string  `json:"dni"`
	Name           string  `json:"name"`
	Age            int     `json:"age"`
	Weight         float64 `json:"weight"`
	Height         float64 `json:"height"`
	Sex            string  `json:"sex"`
	Location       string  `json:"location"`
	CurrentState   string  `json:"current_state"`
	FinalDiagnosis string  `json:"final_diagnosis"`
}

// PatientDTO is used for retrieving a patient
type PatientDTO struct {
	PatientID      uuid.UUID  `json:"patient_id"`
	DNI            string     `json:"dni"`
	Name           string     `json:"name"`
	Age            int        `json:"age"`
	Weight         float64    `json:"weight"`
	Height         float64    `json:"height"`
	Sex            string     `json:"sex"`
	Location       string     `json:"location"`
	CurrentState   string     `json:"current_state"`
	FinalDiagnosis string     `json:"final_diagnosis"`
	LastAlertID    *uuid.UUID `json:"last_alert_id,omitempty"`
	Comorbidities  []string   `json:"comorbidities"`
	Doctors        []string   `json:"doctors"`
	Medications    []string   `json:"medications"`
}

// MapPatientToDTO maps a Patient model to a PatientDTO
func MapPatientToDTO(patient *models.Patient) *PatientDTO {
	var lastAlertId *uuid.UUID
	if patient.Comorbidities == nil {
		patient.Comorbidities = []*models.Comorbidity{}
	}
	if patient.Medications == nil {
		patient.Medications = []*models.Medication{}
	}
	if patient.Alerts != nil {
		lastAlertId = &patient.Alerts[len(patient.Alerts)-1].AlertID
	}
	return &PatientDTO{
		PatientID:      patient.PatientID,
		DNI:            patient.DNI,
		Name:           patient.Name,
		Age:            patient.Age,
		Weight:         patient.Weight,
		Height:         patient.Height,
		Sex:            patient.Sex,
		Location:       patient.Location,
		CurrentState:   patient.CurrentState,
		FinalDiagnosis: patient.FinalDiagnosis,
		LastAlertID:    lastAlertId,
		Comorbidities:  MapComorbiditiesToNames(patient.Comorbidities),
		Medications:    MapMedicationsToMedicationsDetails(patient.Medications),
		Doctors:        MapDoctorsToNames(patient.Doctors),
	}
}

// MapPatientsToDTOs maps a list of Patient models to a list of PatientDTOs
func MapPatientsToDTOs(patients []*models.Patient) []*PatientDTO {
	var patientDTOs []*PatientDTO
	for _, patient := range patients {
		patientDTOs = append(patientDTOs, MapPatientToDTO(patient))
	}
	return patientDTOs
}

// MapCreateDTOToPatient maps a PatientCreateDTO to a Patient model
func MapCreateDTOToPatient(dto *PatientCreateDTO) *models.Patient {
	return &models.Patient{
		DNI:            dto.DNI,
		Name:           dto.Name,
		Age:            dto.Age,
		Weight:         dto.Weight,
		Height:         dto.Height,
		Sex:            dto.Sex,
		Location:       dto.Location,
		CurrentState:   dto.CurrentState,
		FinalDiagnosis: dto.FinalDiagnosis,
	}
}

// MapUpdateDTOToPatient maps a PatientUpdateDTO to a Patient model
func MapUpdateDTOToPatient(dto *PatientUpdateDTO, patient *models.Patient) *models.Patient {
	patient.DNI = dto.DNI
	patient.Name = dto.Name
	patient.Age = dto.Age
	patient.Weight = dto.Weight
	patient.Height = dto.Height
	patient.Sex = dto.Sex
	patient.Location = dto.Location
	patient.CurrentState = dto.CurrentState
	patient.FinalDiagnosis = dto.FinalDiagnosis
	return patient
}
