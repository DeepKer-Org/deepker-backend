package shared

import (
	"gorm.io/gorm"
	"time"
)

// TODO: consider use serial ids for better performance and uuids as external ids
// BaseModel is a struct that contains common fields for all models
type BaseModel struct {
	CreatedAt time.Time      `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-,omitempty" gorm:"index"`
}
