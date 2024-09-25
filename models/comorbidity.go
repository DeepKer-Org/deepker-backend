package models

type Comorbidity struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	PatientID uint   `json:"patient_id"`
	Condition string `gorm:"size:100;not null" json:"comorbidity"`
}
