package models

type Patient struct {
	BaseModel
	PatientID      uint    `gorm:"primaryKey;autoIncrement" json:"patient_id"`
	DNI            string  `gorm:"size:10;unique;not null" json:"dni"`
	Name           string  `gorm:"size:100;not null" json:"name"`
	Age            int     `json:"age"`
	Weight         float64 `gorm:"type:decimal(5,2)" json:"weight"`
	Height         float64 `gorm:"type:decimal(5,2)" json:"height"`
	Sex            string  `gorm:"size:1" json:"sex"`
	Location       string  `gorm:"size:100" json:"location"`
	CurrentState   string  `gorm:"size:50" json:"current_state"`
	FinalDiagnosis string  `gorm:"size:100" json:"final_diagnosis"`
	LastAlertID    *string `gorm:"type:uuid" json:"last_alert_id"`

	Alerts        []*Alert            `json:"alerts"`
	Comorbidities []*Comorbidity      `json:"comorbidities"`
	Medications   []*Medication       `json:"medications"`
	Doctors       []*Doctor           `gorm:"many2many:doctor_patients" json:"doctors"`
	Devices       []*MonitoringDevice `json:"devices"`
}
