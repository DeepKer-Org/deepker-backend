package models

type Medication struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	PatientID  uint   `json:"patient_id"`
	Medication string `gorm:"size:100;not null" json:"medication"`
}
