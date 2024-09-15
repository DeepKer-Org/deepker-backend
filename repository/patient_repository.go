package repository

import (
	"biometric-data-backend/models"
	"github.com/gocql/gocql"
	"log"
	"time"
)

type PatientRepository interface {
	Create(patient *models.Patient) error
	GetAll() ([]*models.Patient, error)
	GetByID(id gocql.UUID) (*models.Patient, error)
	Update(patient *models.Patient) error
	Delete(id gocql.UUID) error
	SoftDelete(id gocql.UUID) error
}

type patientRepository struct {
	session *gocql.Session
}

func NewPatientRepository(session *gocql.Session) PatientRepository {
	return &patientRepository{session: session}
}

// Create a new patient
func (r *patientRepository) Create(patient *models.Patient) error {
	query := `INSERT INTO patients (id, name, age, current_state, medications, created_at) VALUES (?, ?, ?, ?, ?, ?)`
	if err := r.session.Query(query, patient.ID, patient.Name, patient.Age, patient.CurrentState, patient.Medications, patient.Auditable.CreatedAt).Exec(); err != nil {
		log.Println("Error creating patient:", err)
		return err
	}
	return nil
}

// GetAll Get all patients
func (r *patientRepository) GetAll() ([]*models.Patient, error) {
	var patients []*models.Patient
	query := `SELECT id, name, age, current_state, medications, created_at FROM patients`
	iter := r.session.Query(query).Iter()

	for {
		patient := &models.Patient{}
		if !iter.Scan(&patient.ID, &patient.Name, &patient.Age, &patient.CurrentState, &patient.Medications, &patient.Auditable.CreatedAt) {
			break
		}
		patients = append(patients, patient)
	}

	if err := iter.Close(); err != nil {
		return nil, err
	}
	return patients, nil
}

// GetByID Get patient by ID
func (r *patientRepository) GetByID(id gocql.UUID) (*models.Patient, error) {
	patient := &models.Patient{}
	query := `SELECT id, name, age, current_state, medications, created_at FROM patients WHERE id = ? LIMIT 1`
	if err := r.session.Query(query, id).Scan(&patient.ID, &patient.Name, &patient.Age, &patient.CurrentState, &patient.Medications, &patient.Auditable.CreatedAt); err != nil {
		log.Println("Error fetching patient:", err)
		return nil, err
	}
	return patient, nil
}

// Update an existing patient
func (r *patientRepository) Update(patient *models.Patient) error {
	query := `UPDATE patients SET name = ?, age = ?, current_state = ?, medications = ? WHERE id = ?`
	if err := r.session.Query(query, patient.Name, patient.Age, patient.CurrentState, patient.Medications, patient.ID).Exec(); err != nil {
		log.Println("Error updating patient:", err)
		return err
	}
	return nil
}

// Delete a patient by ID
func (r *patientRepository) Delete(id gocql.UUID) error {
	query := `DELETE FROM patients WHERE id = ?`
	if err := r.session.Query(query, id).Exec(); err != nil {
		log.Println("Error deleting patient:", err)
		return err
	}
	return nil
}

// SoftDelete a patient by ID
func (r *patientRepository) SoftDelete(id gocql.UUID) error {
	now := time.Now()
	query := `UPDATE patients SET deleted_at = ? WHERE id = ?`
	if err := r.session.Query(query, now, id).Exec(); err != nil {
		log.Println("Error soft deleting patient:", err)
		return err
	}
	return nil
}
