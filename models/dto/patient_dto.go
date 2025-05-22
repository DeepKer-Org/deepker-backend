package dto

import (
	"github.com/google/uuid"
)

type PatientDTO struct {
	PatientID          uuid.UUID             `json:"patient_id"`
	DNI                string                `json:"dni"`
	Name               string                `json:"name"`
	Age                int                   `json:"age"`
	Weight             float64               `json:"weight"`
	Height             float64               `json:"height"`
	Sex                string                `json:"sex"`
	Location           string                `json:"location"`
	EntryDate          string                `json:"entry_date"`
	MonitoringDeviceID string                `json:"monitoring_device_id"`
	Comorbidities      []string              `json:"comorbidities"`
	MedicalStaff       []*DoctorDTO          `json:"medical_staff"`
	Medications        []*ShortMedicationDTO `json:"medications"`
	MedicalVisits      []*MedicalVisitDTO    `json:"medical_visits"`
}

type PatientFilter struct {
	Name            string
	DNI             string
	Age             int
	DoctorID        string
	Location        string
	DeviceID        string
	ComorbidityName string
	EntryDate       string
	DischargeDate   string
}
