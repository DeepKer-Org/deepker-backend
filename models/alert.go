package models

import (
	"github.com/google/uuid"
	"time"
)

type Alert struct {
	BaseModel
	AttendedTimestamp   *time.Time
	AlertID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	AlertTimestamp      time.Time      `gorm:"not null"`
	AttendedByID        uuid.NullUUID  `gorm:"type:uuid"`
	AttendedBy          *Doctor        `gorm:"foreignKey:AttendedByID"`
	PatientID           uuid.UUID      `gorm:"type:uuid;not null"`
	Patient             *Patient       `gorm:"foreignKey:PatientID;references:PatientID"`
	BiometricDataID     uuid.UUID      `gorm:"type:uuid;not null"`
	BiometricData       *BiometricData `gorm:"foreignKey:BiometricDataID;references:BiometricDataID"`
	Doctors             []*Doctor      `gorm:"many2many:doctor_alerts"`
	ComputerDiagnostics []*ComputerDiagnostic
}
