package models

type Biometrics struct {
	AlertID                string `gorm:"primaryKey;type:uuid" json:"alert_id"`
	O2Saturation           int    `json:"o2_saturation"`
	HeartRate              int    `json:"heart_rate"`
	SystolicBloodPressure  int    `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int    `json:"diastolic_blood_pressure"`
}
