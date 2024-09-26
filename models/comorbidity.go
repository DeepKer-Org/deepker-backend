package models

type Comorbidity struct {
	BaseModel
	ComorbidityID uint   `gorm:"primaryKey;autoIncrement" json:"comorbidity_id"`
	PatientID     uint   `json:"patient_id"`
	Comorbidity   string `gorm:"size:100;not null" json:"comorbidity"`
}
