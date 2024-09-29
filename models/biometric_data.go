package models

import "github.com/google/uuid"

type BiometricData struct {
	BaseModel
	BiometricDataID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	O2Saturation           int       `gorm:"not null"`
	HeartRate              int       `gorm:"not null"`
	SystolicBloodPressure  int       `gorm:"not null"`
	DiastolicBloodPressure int       `gorm:"not null"`
}

func (BiometricData) TableName() string {
	return "biometric_records"
}
