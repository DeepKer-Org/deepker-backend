package models

import (
	"github.com/google/uuid"
)

type ComputerDiagnostic struct {
	BaseModel
	DiagnosticID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Diagnosis    string    `gorm:"size:100;not null"`
	Percentage   float64   `gorm:"type:decimal(4,2)"`
}
