package models

type Biometric struct {
	BaseModel
	BiometricsID           string `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"biometrics_id"`
	AlertID                string `gorm:"type:uuid;not null" json:"alert_id"`
	O2Saturation           int    `json:"o2_saturation"`
	HeartRate              int    `json:"heart_rate"`
	SystolicBloodPressure  int    `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int    `json:"diastolic_blood_pressure"`
}
