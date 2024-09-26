package models

import (
	"time"
)

type Patient struct {
	ID             uint       `gorm:"primaryKey" json:"patient_id"`
	DNI            string     `gorm:"size:10;unique;not null" json:"dni"`
	Name           string     `gorm:"size:100;not null" json:"name"`
	Age            int        `json:"age"`
	Weight         float64    `json:"weight"`
	Height         float64    `json:"height"`
	Sex            string     `gorm:"size:1" json:"sex"`
	Location       string     `gorm:"size:100" json:"location"`
	CurrentState   string     `gorm:"size:50" json:"current_state"`
	FinalDiagnosis string     `gorm:"size:100" json:"final_diagnosis"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DeletedAt      *time.Time `json:"deleted_at,omitempty"`

	Comorbidities []Comorbidity   `gorm:"foreignKey:PatientID" json:"comorbidities"`
	Medications   []Medication    `gorm:"foreignKey:PatientID" json:"medications"`
	MedicalStaff  []PatientDoctor `gorm:"foreignKey:PatientID" json:"medical_staff"`
}
