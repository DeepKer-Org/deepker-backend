package dtos

type CreateAlertDTO struct {
	AlertID           string                 `json:"alert_id" binding:"required"`
	Status            string                 `json:"alert_status" binding:"required"`
	AttendedBy        string                 `json:"attended_by"`
	AlertTimestamp    string                 `json:"alert_timestamp" binding:"required"`
	AttendedTimestamp string                 `json:"attended_timestamp"`
	PatientID         uint                   `json:"patient_id" binding:"required"`
	Biometrics        BiometricsDTO          `json:"biometrics" binding:"required"`
	ComputerDiagnoses []ComputerDiagnosisDTO `json:"computer_diagnoses"`
	AssociatedDoctors []string               `json:"associated_doctors"`
}

type BiometricsDTO struct {
	O2Saturation           int `json:"o2_saturation"`
	HeartRate              int `json:"heart_rate"`
	SystolicBloodPressure  int `json:"systolic_blood_pressure"`
	DiastolicBloodPressure int `json:"diastolic_blood_pressure"`
}

type ComputerDiagnosisDTO struct {
	Diagnosis  string  `json:"diagnosis"`
	Percentage float64 `json:"percentage"`
}
