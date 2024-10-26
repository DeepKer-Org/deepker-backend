package models

import "github.com/google/uuid"

type Phone struct {
	BaseModel
	PhoneID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ExponentPushToken string    `gorm:"type:varchar(255);not null"`
}
