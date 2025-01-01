package comorbidity

import (
	"biometric-data-backend/internal/shared"
	"github.com/google/uuid"
)

type Model struct {
	shared.BaseModel
	ComorbidityID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	PatientID     uuid.UUID `gorm:"type:uuid"`
	Comorbidity   string    `gorm:"size:100;not null"`
}

func (Model) TableName() string {
	return "comorbidities"
}
