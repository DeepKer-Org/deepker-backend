package models

type ComputerDiagnosis struct {
	BaseModel
	DiagnosisID uint    `gorm:"primaryKey;autoIncrement" json:"diagnosis_id"`
	AlertID     string  `gorm:"type:uuid;not null" json:"alert_id"`
	Diagnosis   string  `gorm:"size:100;not null" json:"diagnosis"`
	Percentage  float64 `gorm:"type:decimal(4,2)" json:"percentage"`
}
