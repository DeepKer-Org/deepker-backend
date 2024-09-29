package models

import (
	"github.com/google/uuid"
)

type Doctor struct {
	BaseModel
	DoctorID       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"doctor_id"`
	DNI            string     `gorm:"size:10;unique;not null" json:"dni"`
	Name           string     `gorm:"size:100;not null" json:"name"`
	Password       string     `gorm:"size:100;not null" json:"password"`
	Specialization string     `gorm:"size:100" json:"specialization"`
	Alerts         []*Alert   `gorm:"many2many:doctor_alerts" json:"alerts"`
	Patients       []*Patient `gorm:"many2many:doctor_patients" json:"patients"`
}
