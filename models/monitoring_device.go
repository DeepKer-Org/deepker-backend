package models

import (
	"github.com/google/uuid"
)

type MonitoringDevice struct {
	BaseModel
	DeviceID   string     `gorm:"size:10;primaryKey" json:"device_id"`
	Status     string     `gorm:"size:50;not null;check:status in ('In Use', 'Free', 'Unavailable', 'Connecting')" json:"status"`
	PatientID  *uuid.UUID `gorm:"type:uuid" json:"patient_id"`
	Patient    *Patient   `gorm:"foreignKey:PatientID;references:PatientID"`
	LinkedByID *uuid.UUID `gorm:"type:uuid" json:"linked_by_id"`
	LinkedBy   *Doctor    `gorm:"foreignKey:LinkedByID"`
}
