package models

import "time"

type Alert struct {
	BaseModel
	AlertID             string                `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"alert_id"`
	AlertStatus         string                `gorm:"size:50;not null" json:"alert_status"`
	Room                string                `gorm:"size:100" json:"room"`
	AlertTimestamp      time.Time             `gorm:"not null" json:"alert_timestamp"`
	AttendedTimestamp   *time.Time            `json:"attended_timestamp"`
	PatientID           uint                  `gorm:"not null" json:"patient_id"`
	Biometrics          []*Biometric          `json:"biometrics"`
	ComputerDiagnostics []*ComputerDiagnostic `json:"computer_diagnostics"`
	Doctors             []*Doctor             `gorm:"many2many:doctor_alerts" json:"doctors"`
}
