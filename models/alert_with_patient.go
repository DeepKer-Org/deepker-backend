package models

// AlertWithPatient combines alert details with the patient information
type AlertWithPatient struct {
	Alert   Alert   `json:"alert"`
	Patient Patient `json:"patient"`
}
