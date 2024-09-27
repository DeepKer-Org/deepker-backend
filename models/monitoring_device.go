package models

import (
	"github.com/google/uuid"
)

type MonitoringDevice struct {
	BaseModel
	DeviceID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"device_id"`
	Type      string    `gorm:"size:50;not null" json:"type"`
	Status    string    `gorm:"size:50;not null;check:status in ('In Use', 'Free', 'Unavailable')" json:"status"`
	PatientID uuid.UUID `gorm:"type:uuid" json:"patient_id"`
	Sensors   []string  `gorm:"type:text[]" json:"sensors"`
}
