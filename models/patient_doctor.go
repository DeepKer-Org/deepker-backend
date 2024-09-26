package models

type PatientDoctor struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	PatientID uint   `json:"patient_id"`
	Doctor    string `gorm:"size:100;not null" json:"doctor_name"`
}
