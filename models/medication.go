package models

import "time"

type Medication struct {
	BaseModel
	MedicationID uint       `gorm:"primaryKey;autoIncrement" json:"medication_id"`
	PatientID    uint       `json:"patient_id"`
	Medication   string     `gorm:"size:100;not null" json:"medication"`
	StartDate    *time.Time `json:"start_date"`
	EndDate      *time.Time `json:"end_date"`
	Dosage       string     `gorm:"size:50" json:"dosage"`
	Periodicity  string     `gorm:"size:50" json:"periodicity"`
}
