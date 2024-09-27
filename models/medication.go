package models

import (
	"github.com/google/uuid"
	"time"
)

type Medication struct {
	BaseModel
	MedicationID uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"medication_id"`
	PatientID    uuid.UUID  `gorm:"type:uuid;not null" json:"patient_id"`
	Medication   string     `gorm:"size:100;not null" json:"medication"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Dosage       string     `gorm:"size:50" json:"dosage"`
	Periodicity  string     `gorm:"size:50" json:"periodicity"`
}
