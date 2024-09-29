package models

import (
	"github.com/google/uuid"
	"time"
)

type Alert struct {
	BaseModel
	AlertID             uuid.UUID             `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"alert_id"`
	AlertStatus         string                `gorm:"size:50;not null" json:"alert_status"`
	Room                string                `gorm:"size:100" json:"room"`
	AlertTimestamp      time.Time             `gorm:"not null" json:"alert_timestamp"`
	AttendedTimestamp   *time.Time            `json:"attended_timestamp"`
	AttendedBy          uuid.UUID             `gorm:"type:uuid" json:"attended_by"`
	PatientID           uuid.UUID             `gorm:"type:uuid;not null" json:"patient_id"`
	Biometrics          []*Biometric          `json:"biometrics"`
	ComputerDiagnostics []*ComputerDiagnostic `json:"computer_diagnostics"`
	Doctors             []*Doctor             `gorm:"many2many:doctor_alerts" json:"doctors"`
}
