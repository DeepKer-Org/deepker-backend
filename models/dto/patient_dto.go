package dto

import (
	"github.com/google/uuid"
)

type PatientDTO struct {
	PatientID      uuid.UUID  `json:"patient_id"`
	DNI            string     `json:"dni"`
	Name           string     `json:"name"`
	Age            int        `json:"age"`
	Weight         float64    `json:"weight"`
	Height         float64    `json:"height"`
	Sex            string     `json:"sex"`
	Location       string     `json:"location,omitempty"`
	CurrentState   string     `json:"current_state,omitempty"`
	FinalDiagnosis string     `json:"final_diagnosis,omitempty"`
	LastAlertID    *uuid.UUID `json:"last_alert_id,omitempty"`
	Comorbidities  []string   `json:"comorbidities,omitempty"`
	Doctors        []string   `json:"doctors,omitempty"`
	Medications    []string   `json:"medications,omitempty"`
}
