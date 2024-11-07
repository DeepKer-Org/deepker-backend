package models

import "github.com/google/uuid"

type BiometricData struct {
	BaseModel
	BiometricDataID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	O2Saturation    float64   `gorm:"not null"`
	HeartRate       float64   `gorm:"not null"`
}

func (BiometricData) TableName() string {
	return "biometric_data"
}
