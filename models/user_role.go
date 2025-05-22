package models

import "github.com/google/uuid"

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();foreignKey"`
	RoleID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();foreignKey"`
}
