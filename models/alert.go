package models

import "github.com/gocql/gocql"

// BloodPressure structure for systolic and diastolic values
type BloodPressure struct {
	Systolic  int `json:"systolic"`
	Diastolic int `json:"diastolic"`
}

// Alert represents an alert for a patient's vital signs
type Alert struct {
	AlertID           gocql.UUID    `json:"alertId"`
	PatientID         gocql.UUID    `json:"patientId"`
	Room              string        `json:"room"`
	AlertTimestamp    int64         `json:"alertTimestamp"`
	O2Saturation      int           `json:"O2Saturation"`
	HeartRate         int           `json:"heartRate"`
	BloodPressure     BloodPressure `json:"bloodPressure"`
	ComputerDiagnoses []string      `json:"computerDiagnoses"`
	AlertStatus       string        `json:"alertStatus"`
	AttendedBy        string        `json:"attendedBy"`
	AttendedTimestamp int64         `json:"attendedTimestamp"`
	FinalDiagnosis    string        `json:"finalDiagnosis"`
	Auditable         Auditable     `json:"auditable"`
}
