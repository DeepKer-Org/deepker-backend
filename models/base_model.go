package models

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-,omitempty" gorm:"index"`
}
