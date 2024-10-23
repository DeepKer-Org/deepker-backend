package models

import (
	"biometric-data-backend/enums"
	"github.com/google/uuid"
)

type Role struct {
	RoleID   uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleName enums.RoleEnum `gorm:"size:50;unique;not null"`
}
