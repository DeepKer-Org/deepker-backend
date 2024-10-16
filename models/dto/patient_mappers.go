package dto

import (
	"biometric-data-backend/models"
)

func MapPatientToDTO(patient *models.Patient) *PatientDTO {
	lastAlertId := ""
	if len(patient.Alerts) > 0 {
		lastAlertId = patient.Alerts[len(patient.Alerts)-1].AlertID.String()
	}

	return &PatientDTO{
		PatientID:     patient.PatientID,
		DNI:           patient.DNI,
		Name:          patient.Name,
		Age:           patient.Age,
		Weight:        patient.Weight,
		Height:        patient.Height,
		Sex:           patient.Sex,
		Location:      patient.Location,
		CurrentState:  patient.CurrentState,
		Comorbidities: MapComorbiditiesToNames(patient.Comorbidities),
		Medications:   MapMedicationsToMedicationsDetails(patient.Medications),
		Doctors:       MapDoctorsToNames(patient.Doctors),
		LastAlertID:   lastAlertId,
	}
}

func MapPatientsToDTOs(patients []*models.Patient) []*PatientDTO {
	var patientDTOs []*PatientDTO
	for _, patient := range patients {
		patientDTOs = append(patientDTOs, MapPatientToDTO(patient))
	}
	return patientDTOs
}

func MapPatientToPatientForAlertDTO(patient *models.Patient) *PatientForAlertDTO {
	if patient.Medications == nil {
		patient.Medications = make([]*models.Medication, 0)
	}

	return &PatientForAlertDTO{
		DNI:           patient.DNI,
		Name:          patient.Name,
		Location:      patient.Location,
		Age:           patient.Age,
		Sex:           patient.Sex,
		Doctors:       MapDoctorsToNames(patient.Doctors),
		Comorbidities: MapComorbiditiesToNames(patient.Comorbidities),
		Medications:   MapShortMedicationsToDTOs(patient.Medications),
	}
}

func MapCreateDTOToPatient(dto *PatientCreateDTO) *models.Patient {
	return &models.Patient{
		DNI:          dto.DNI,
		Name:         dto.Name,
		Age:          dto.Age,
		Weight:       dto.Weight,
		Height:       dto.Height,
		Sex:          dto.Sex,
		Location:     dto.Location,
		CurrentState: dto.CurrentState,
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
	return patient
}
