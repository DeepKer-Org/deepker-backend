package dto

type PatientForAlertDTO struct {
	DNI                string                `json:"dni"`
	Name               string                `json:"name"`
	Location           string                `json:"location"`
	Age                int                   `json:"age"`
	Sex                string                `json:"sex"`
	MonitoringDeviceID string                `json:"monitoring_device_id"`
	Comorbidities      []string              `json:"comorbidities"`
	Medications        []*ShortMedicationDTO `json:"medications"`
}
