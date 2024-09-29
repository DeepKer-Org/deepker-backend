package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

// PatientCreateDTO is used for creating a new patient
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
	DNI            string     `json:"dni"`
	Name           string     `json:"name"`
	Age            int        `json:"age"`
	Weight         float64    `json:"weight"`
	Height         float64    `json:"height"`
	Sex            string     `json:"sex"`
	Location       string     `json:"location"`
	CurrentState   string     `json:"current_state"`
	FinalDiagnosis string     `json:"final_diagnosis"`
	LastAlertID    *uuid.UUID `json:"last_alert_id"`
}

// PatientDTO is used for retrieving a patient along with related entities
type PatientDTO struct {
	PatientID      uuid.UUID              `json:"patient_id"`
	DNI            string                 `json:"dni"`
	Name           string                 `json:"name"`
	Age            int                    `json:"age"`
	Weight         float64                `json:"weight"`
	Height         float64                `json:"height"`
	Sex            string                 `json:"sex"`
	Location       string                 `json:"location"`
	CurrentState   string                 `json:"current_state"`
	FinalDiagnosis string                 `json:"final_diagnosis"`
	LastAlertID    *uuid.UUID             `json:"last_alert_id"`
	Alerts         []*AlertDTO            `json:"alerts"`
	Comorbidities  []*ComorbidityDTO      `json:"comorbidities"`
	Medications    []*MedicationDTO       `json:"medications"`
	Doctors        []*DoctorDTO           `json:"doctors"`
	Devices        []*MonitoringDeviceDTO `json:"devices"`
}

// MapPatientToDTO maps a Patient model to a PatientDTO
func MapPatientToDTO(patient *models.Patient) *PatientDTO {
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
		LastAlertID:    patient.LastAlertID,
		Alerts:         MapAlertsToDTOs(patient.Alerts),
		Comorbidities:  MapComorbiditiesToDTOs(patient.Comorbidities),
		Medications:    MapMedicationsToDTOs(patient.Medications),
		Doctors:        MapDoctorsToDTOs(patient.Doctors),
		Devices:        MapMonitoringDevicesToDTOs(patient.Devices),
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
