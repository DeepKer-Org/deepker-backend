package dto

import (
	"github.com/google/uuid"
)

type PatientDTO struct {
	PatientID     uuid.UUID `json:"patient_id"`
	DNI           string    `json:"dni"`
	Name          string    `json:"name"`
	Age           int       `json:"age"`
	Weight        float64   `json:"weight"`
	Height        float64   `json:"height"`
	Sex           string    `json:"sex"`
	Location      string    `json:"location"`
	CurrentState  string    `json:"current_state"`
	Comorbidities []string  `json:"comorbidities"`
	Doctors       []string  `json:"doctors"`
	Medications   []string  `json:"medications"`
	LastAlertID   string    `json:"last_alert_id"`
}
