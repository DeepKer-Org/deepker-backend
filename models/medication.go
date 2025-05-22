package models

import (
	"github.com/google/uuid"
	"time"
)

type Medication struct {
	BaseModel
	MedicationID uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientID    uuid.UUID  `gorm:"type:uuid;not null"`
	Name         string     `gorm:"size:100;not null"`
	StartDate    *time.Time `gorm:"type:date"`
	EndDate      *time.Time `gorm:"type:date"`
	Dosage       string     `gorm:"size:50"`
	Periodicity  string     `gorm:"size:50"`
}
