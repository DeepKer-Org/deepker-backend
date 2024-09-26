package models

type MonitoringDevice struct {
	BaseModel
	DeviceID  string   `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"device_id"`
	Type      string   `gorm:"size:50;not null" json:"type"`
	Status    string   `gorm:"size:50;not null;check:status in ('In Use', 'Free', 'Unavailable')" json:"status"`
	PatientID uint     `json:"patient_id"`
	Sensors   []string `gorm:"type:text[]" json:"sensors"`
}
