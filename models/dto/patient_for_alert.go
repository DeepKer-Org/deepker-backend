package dto

type PatientForAlertDTO struct {
	DNI            string   `json:"dni"`
	Name           string   `json:"name"`
	Location       string   `json:"current_location,omitempty"`
	FinalDiagnosis string   `json:"final_diagnosis,omitempty"`
	Doctors        []string `json:"doctors,omitempty"`
}
