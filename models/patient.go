package models

import (
	"github.com/google/uuid"
)

type Patient struct {
	BaseModel
	PatientID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	DNI           string    `gorm:"size:10;unique;not null"`
	Name          string    `gorm:"size:100;not null"`
	Age           int       `gorm:"not null"`
	Weight        float64   `gorm:"type:decimal(5,2);not null"`
	Height        float64   `gorm:"type:decimal(5,2);not null"`
	Sex           string    `gorm:"size:1;not null"`
	Location      string    `gorm:"size:100;default:null"`
	CurrentState  string    `gorm:"size:50;default:null"`
	Alerts        []*Alert
	Comorbidities []*Comorbidity `gorm:"foreignKey:PatientID;references:PatientID"`
	Medications   []*Medication  `gorm:"foreignKey:PatientID;references:PatientID"`
	Devices       []*MonitoringDevice
	Doctors       []*Doctor `gorm:"many2many:doctor_patients;foreignKey:PatientID;joinForeignKey:PatientID;References:DoctorID;joinReferences:DoctorID"`
}
