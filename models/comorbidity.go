package models

import (
	"github.com/google/uuid"
)

type Comorbidity struct {
	BaseModel
	ComorbidityID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	PatientID     uuid.UUID `gorm:"type:uuid"`
	Comorbidity   string    `gorm:"size:100;not null"`
}
