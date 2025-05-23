package dto

import (
	"biometric-data-backend/models"
)

func MapPatientToDTO(patient *models.Patient) *PatientDTO {
	var monitoringDeviceID string
	if patient.MonitoringDevice == nil {
		monitoringDeviceID = ""
	} else {
		monitoringDeviceID = patient.MonitoringDevice.DeviceID
	}
	return &PatientDTO{
		PatientID:          patient.PatientID,
		DNI:                patient.DNI,
		Name:               patient.Name,
		Age:                patient.Age,
		Weight:             patient.Weight,
		Height:             patient.Height,
		Sex:                patient.Sex,
		Location:           patient.Location,
		MonitoringDeviceID: monitoringDeviceID,
		EntryDate:          FindEntryDate(patient.MedicalVisits),
		Comorbidities:      MapComorbiditiesToNames(patient.Comorbidities),
		MedicalStaff:       MapDoctorsToDTOs(patient.Doctors),
		Medications:        MapShortMedicationsToDTOs(patient.Medications),
		MedicalVisits:      MapMedicalVisitsToDTOs(patient.MedicalVisits),
	}
}

func MapPatientsToDTOs(patients []*models.Patient) []*PatientDTO {
	var patientDTOs = make([]*PatientDTO, 0)
	for _, patient := range patients {
		patientDTOs = append(patientDTOs, MapPatientToDTO(patient))
	}
	return patientDTOs
}

func MapPatientToPatientForAlertDTO(patient *models.Patient) *PatientForAlertDTO {
	if patient.Medications == nil {
		patient.Medications = make([]*models.Medication, 0)
	}
	var monitoringDeviceID string
	if patient.MonitoringDevice == nil {
		monitoringDeviceID = ""
	} else {
		monitoringDeviceID = patient.MonitoringDevice.DeviceID
	}

	return &PatientForAlertDTO{
		DNI:                patient.DNI,
		Name:               patient.Name,
		Location:           patient.Location,
		Age:                patient.Age,
		Sex:                patient.Sex,
		MonitoringDeviceID: monitoringDeviceID,
		Comorbidities:      MapComorbiditiesToNames(patient.Comorbidities),
		Medications:        MapShortMedicationsToDTOs(patient.Medications),
	}
}

func MapPatientToPatientForDeviceDTO(patient *models.Patient) *PatientForDeviceDTO {

	return &PatientForDeviceDTO{
		PatientID: patient.PatientID,
		DNI:       patient.DNI,
		Name:      patient.Name,
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
