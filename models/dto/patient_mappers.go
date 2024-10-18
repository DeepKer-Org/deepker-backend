package dto

import (
	"biometric-data-backend/models"
)

func MapPatientToDTO(patient *models.Patient) *PatientDTO {
	return &PatientDTO{
		PatientID:     patient.PatientID,
		DNI:           patient.DNI,
		Name:          patient.Name,
		Age:           patient.Age,
		Weight:        patient.Weight,
		Height:        patient.Height,
		Sex:           patient.Sex,
		Location:      patient.Location,
		EntryDate:     FindEntryDate(patient.MedicalVisits),
		Comorbidities: MapComorbiditiesToNames(patient.Comorbidities),
		MedicalStaff:  MapDoctorsToDTOs(patient.Doctors),
		Medications:   MapShortMedicationsToDTOs(patient.Medications),
		MedicalVisits: MapMedicalVisitsToDTOs(patient.MedicalVisits),
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
		Comorbidities: MapComorbiditiesToNames(patient.Comorbidities),
		Medications:   MapShortMedicationsToDTOs(patient.Medications),
	}
}

func MapCreateDTOToPatient(dto *PatientCreateDTO) *models.Patient {
	return &models.Patient{
		DNI:      dto.DNI,
		Name:     dto.Name,
		Age:      dto.Age,
		Weight:   dto.Weight,
		Height:   dto.Height,
		Sex:      dto.Sex,
		Location: dto.Location,
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
	return patient
}
