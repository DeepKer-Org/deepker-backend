package dto

import "github.com/google/uuid"

// UpdateDTO is used for updating an existing comorbidity
type UpdateDTO struct {
	PatientID   uuid.UUID `json:"patient_id"`
	Comorbidity string    `json:"comorbidity"`
}
