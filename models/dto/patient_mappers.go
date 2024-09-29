package dto

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
)

func MapPatientToDTO(patient *models.Patient) *PatientDTO {
	var lastAlertId *uuid.UUID
	if len(patient.Alerts) > 0 {
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

func MapPatientsToDTOs(patients []*models.Patient) []*PatientDTO {
	var patientDTOs []*PatientDTO
	for _, patient := range patients {
		patientDTOs = append(patientDTOs, MapPatientToDTO(patient))
	}
	return patientDTOs
}

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
