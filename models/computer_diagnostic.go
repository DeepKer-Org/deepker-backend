package models

import (
	"github.com/google/uuid"
)

type ComputerDiagnostic struct {
	BaseModel
	DiagnosisID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"diagnosis_id"`
	AlertID     uuid.UUID `gorm:"type:uuid;not null" json:"alert_id"`
	Diagnosis   string    `gorm:"size:100;not null" json:"diagnosis"`
	Percentage  float64   `gorm:"type:decimal(4,2)" json:"percentage"`
}
