package dto

import "github.com/google/uuid"

// ResponseDTO is used for retrieving a comorbidity
type ResponseDTO struct {
	ComorbidityID uuid.UUID `json:"comorbidity_id"`
	PatientID     uuid.UUID `json:"patient_id"`
	Comorbidity   string    `json:"comorbidity"`
}
