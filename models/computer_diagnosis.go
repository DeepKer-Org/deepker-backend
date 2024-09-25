package models

type ComputerDiagnosis struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
	AlertID    string  `json:"alert_id"`
	Diagnosis  string  `gorm:"size:100;not null" json:"diagnosis"`
	Percentage float64 `gorm:"type:decimal(4,2)" json:"percentage"`
}
