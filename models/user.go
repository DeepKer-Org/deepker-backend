package models

import (
	"github.com/google/uuid"
)

type User struct {
	UserID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string    `gorm:"size:100;unique;not null"`
	Password string    `gorm:"not null"`
	Roles    []*Role   `gorm:"many2many:user_roles;joinForeignKey:UserID;joinReferences:RoleID;"`
}
