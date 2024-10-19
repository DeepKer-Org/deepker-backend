package models

import (
	"github.com/google/uuid"
	"time"
)

type MedicalVisit struct {
	BaseModel
	MedicalVisitID uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"medical_visit_id"`
	PatientID      uuid.UUID  `gorm:"type:uuid;not null"`
	Reason         string     `gorm:"size:100;not null" json:"reason"`
	Diagnosis      string     `gorm:"size:100;not null" json:"diagnosis"`
	Treatment      string     `gorm:"size:100" json:"treatment"`
	EntryDate      *time.Time `json:"entry_date" gorm:"autoCreateTime"`
	DischargeDate  *time.Time `json:"discharge_date" gorm:"type:date"`
}
