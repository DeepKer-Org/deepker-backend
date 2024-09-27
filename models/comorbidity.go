package models

import (
	"github.com/google/uuid"
)

type Comorbidity struct {
	BaseModel
	ComorbidityID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"comorbidity_id"`
	PatientID     uuid.UUID `gorm:"type:uuid" json:"patient_id"`
	Comorbidity   string    `gorm:"size:100;not null" json:"comorbidity"`
}
