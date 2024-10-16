package dto

type PatientForAlertDTO struct {
	DNI            string                `json:"dni"`
	Name           string                `json:"name"`
	Location       string                `json:"location"`
	Age            int                   `json:"age"`
	Sex            string                `json:"sex"`
	FinalDiagnosis string                `json:"final_diagnosis"`
	Doctors        []string              `json:"doctors"`
	Comorbidities  []string              `json:"comorbidities"`
	Medications    []*ShortMedicationDTO `json:"medications"`
}
