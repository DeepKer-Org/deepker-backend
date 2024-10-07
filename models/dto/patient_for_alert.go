package dto

type PatientForAlertDTO struct {
	DNI            string   `json:"dni"`
	Name           string   `json:"name"`
	Location       string   `json:"current_location"`
	FinalDiagnosis string   `json:"final_diagnosis"`
	Doctors        []string `json:"doctors"`
}
