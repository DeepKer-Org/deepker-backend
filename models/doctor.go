package models

import (
	"github.com/google/uuid"
	"time"
)

type Doctor struct {
	BaseModel
	DoctorID       uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"doctor_id"`
	DNI            string     `gorm:"size:10;unique;not null" json:"dni"`
	IssuanceDate   time.Time  `gorm:"not null" json:"issuance_date" time_format:"2006-01-02"`
	Name           string     `gorm:"size:100;not null" json:"name"`
	Specialization string     `gorm:"size:100" json:"specialization"`
	UserID         uuid.UUID  `gorm:"type:uuid;not null; unique"`
	User           User       `gorm:"foreignKey:UserID;references:UserID"`
	Alerts         []*Alert   `gorm:"many2many:doctor_alerts" json:"alerts"`
	Patients       []*Patient `gorm:"many2many:doctor_patients" json:"patients"`
	AttendedAlerts []*Alert   `gorm:"foreignKey:AttendedByID" json:"attended_alerts"`
}
