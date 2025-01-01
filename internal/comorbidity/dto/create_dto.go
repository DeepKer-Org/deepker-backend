package dto

import "github.com/google/uuid"

// CreateDTO is used for the creation of a new comorbidity
type CreateDTO struct {
	PatientID   uuid.UUID `json:"patient_id"`
	Comorbidity string    `json:"comorbidity"`
}
