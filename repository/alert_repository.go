package repository

import (
	"biometric-data-backend/models"
	"github.com/gocql/gocql"
	"log"
)

type AlertRepository interface {
	Create(alert *models.Alert) error
	GetAll() ([]*models.Alert, error)
	GetByID(id gocql.UUID) (*models.Alert, error)
	Update(alert *models.Alert) error
	Delete(id gocql.UUID) error
	SoftDelete(id gocql.UUID) error
	GetAlertWithPatient(id gocql.UUID) (*models.AlertWithPatient, error) // Método para obtener alerta con datos del paciente
}

type alertRepository struct {
	session *gocql.Session
}

func NewAlertRepository(session *gocql.Session) AlertRepository {
	return &alertRepository{session: session}
}

// Create a new alert
func (r *alertRepository) Create(alert *models.Alert) error {
	query := `INSERT INTO alerts (alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	if err := r.session.Query(query, alert.AlertID, alert.PatientID, alert.Room, alert.AlertTimestamp, alert.O2Saturation, alert.HeartRate, map[string]int{
		"systolic":  alert.BloodPressure.Systolic,
		"diastolic": alert.BloodPressure.Diastolic,
	}, alert.ComputerDiagnoses, alert.AlertStatus).Exec(); err != nil {
		log.Println("Error creating alert:", err)
		return err
	}
	return nil
}

// GetAll Get all alerts
func (r *alertRepository) GetAll() ([]*models.Alert, error) {
	var alerts []*models.Alert
	query := `SELECT alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status FROM alerts`
	iter := r.session.Query(query).Iter()

	for {
		alert := &models.Alert{}
		bloodPressure := make(map[string]int)
		if !iter.Scan(&alert.AlertID, &alert.PatientID, &alert.Room, &alert.AlertTimestamp, &alert.O2Saturation, &alert.HeartRate, &bloodPressure, &alert.ComputerDiagnoses, &alert.AlertStatus) {
			break
		}
		alert.BloodPressure.Systolic = bloodPressure["systolic"]
		alert.BloodPressure.Diastolic = bloodPressure["diastolic"]
		alerts = append(alerts, alert)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return alerts, nil
}

// GetByID Get alert by ID
func (r *alertRepository) GetByID(id gocql.UUID) (*models.Alert, error) {
	alert := &models.Alert{}
	bloodPressure := make(map[string]int)
	query := `SELECT alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status FROM alerts WHERE alert_id = ? LIMIT 1`
	if err := r.session.Query(query, id).Scan(&alert.AlertID, &alert.PatientID, &alert.Room, &alert.AlertTimestamp, &alert.O2Saturation, &alert.HeartRate, &bloodPressure, &alert.ComputerDiagnoses, &alert.AlertStatus); err != nil {
		log.Println("Error fetching alert:", err)
		return nil, err
	}
	alert.BloodPressure.Systolic = bloodPressure["systolic"]
	alert.BloodPressure.Diastolic = bloodPressure["diastolic"]
	return alert, nil
}

// Update an existing alert
func (r *alertRepository) Update(alert *models.Alert) error {
	query := `UPDATE alerts SET room = ?, o2_saturation = ?, heart_rate = ?, blood_pressure = ?, computer_diagnoses = ?, alert_status = ? WHERE alert_id = ?`
	if err := r.session.Query(query, alert.Room, alert.O2Saturation, alert.HeartRate, map[string]int{
		"systolic":  alert.BloodPressure.Systolic,
		"diastolic": alert.BloodPressure.Diastolic,
	}, alert.ComputerDiagnoses, alert.AlertStatus, alert.AlertID).Exec(); err != nil {
		log.Println("Error updating alert:", err)
		return err
	}
	return nil
}

// Delete an alert by ID
func (r *alertRepository) Delete(id gocql.UUID) error {
	query := `DELETE FROM alerts WHERE alert_id = ?`
	if err := r.session.Query(query, id).Exec(); err != nil {
		log.Println("Error deleting alert:", err)
		return err
	}
	return nil
}

// SoftDelete an alert by ID
func (r *alertRepository) SoftDelete(id gocql.UUID) error {
	query := `UPDATE alerts SET alert_status = 'DELETED' WHERE alert_id = ?`
	if err := r.session.Query(query, id).Exec(); err != nil {
		log.Println("Error soft deleting alert:", err)
		return err
	}
	return nil
}

// GetAlertWithPatient Get an alert along with the associated patient details
func (r *alertRepository) GetAlertWithPatient(id gocql.UUID) (*models.AlertWithPatient, error) {
	alertWithPatient := &models.AlertWithPatient{}
	bloodPressure := make(map[string]int)

	// Query to get the alert details
	queryAlert := `SELECT alert_id, patient_id, room, alert_timestamp, o2_saturation, heart_rate, blood_pressure, computer_diagnoses, alert_status 
				   FROM alerts WHERE alert_id = ? LIMIT 1`
	if err := r.session.Query(queryAlert, id).Scan(&alertWithPatient.Alert.AlertID, &alertWithPatient.Alert.PatientID, &alertWithPatient.Alert.Room,
		&alertWithPatient.Alert.AlertTimestamp, &alertWithPatient.Alert.O2Saturation, &alertWithPatient.Alert.HeartRate, &bloodPressure,
		&alertWithPatient.Alert.ComputerDiagnoses, &alertWithPatient.Alert.AlertStatus); err != nil {
		log.Println("Error fetching alert:", err)
		return nil, err
	}

	alertWithPatient.Alert.BloodPressure.Systolic = bloodPressure["systolic"]
	alertWithPatient.Alert.BloodPressure.Diastolic = bloodPressure["diastolic"]

	// Query to get the patient details associated with the alert
	queryPatient := `SELECT id, name, age, current_state, medications, created_at FROM patients WHERE id = ? LIMIT 1`
	if err := r.session.Query(queryPatient, alertWithPatient.Alert.PatientID).Scan(&alertWithPatient.Patient.ID, &alertWithPatient.Patient.Name,
		&alertWithPatient.Patient.Age, &alertWithPatient.Patient.CurrentState, &alertWithPatient.Patient.Medications, &alertWithPatient.Patient.Auditable.CreatedAt); err != nil {
		log.Println("Error fetching patient:", err)
		return nil, err
	}

	return alertWithPatient, nil
}
