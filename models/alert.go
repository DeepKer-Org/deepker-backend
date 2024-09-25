package models

import (
	"gorm.io/gorm"
	"time"
)

type Alert struct {
	AlertID           string              `gorm:"primaryKey;type:uuid" json:"alert_id"`
	AlertStatus       string              `gorm:"size:50;not null" json:"alert_status"`
	AttendedBy        string              `gorm:"size:100" json:"attended_by"`
	AlertTimestamp    time.Time           `json:"alert_timestamp"`
	AttendedTimestamp *time.Time          `json:"attended_timestamp,omitempty"`
	PatientID         uint                `json:"patient_id"`
	CreatedAt         time.Time           `json:"created_at"`
	UpdatedAt         time.Time           `json:"updated_at"`
	DeletedAt         gorm.DeletedAt      `gorm:"index"`
	Biometrics        Biometrics          `gorm:"foreignKey:AlertID" json:"biometrics"`
	ComputerDiagnoses []ComputerDiagnosis `gorm:"foreignKey:AlertID" json:"computer_diagnoses"`
	AssociatedDoctors []AlertDoctor       `gorm:"foreignKey:AlertID" json:"associated_doctors"`
}
