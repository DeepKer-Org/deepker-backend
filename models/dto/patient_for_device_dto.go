package dto

import "github.com/google/uuid"

type PatientForDeviceDTO struct {
	PatientID uuid.UUID `json:"patient_id"`
	DNI       string    `json:"dni"`
	Name      string    `json:"name"`
}
