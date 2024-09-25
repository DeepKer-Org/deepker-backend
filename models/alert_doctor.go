package models

type AlertDoctor struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	AlertID string `json:"alert_id"`
	Doctor  string `gorm:"size:100;not null" json:"doctor_name"`
}
