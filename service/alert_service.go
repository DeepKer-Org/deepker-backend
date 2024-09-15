package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/repository"
	"github.com/gocql/gocql"
	"time"
)

type AlertService interface {
	CreateAlert(alert *models.Alert) error
	GetAllAlerts() ([]*models.Alert, error)
	GetAlertByID(id gocql.UUID) (*models.Alert, error)
	UpdateAlert(alert *models.Alert) error
	DeleteAlert(id gocql.UUID) error
	GetAlertWithPatient(id gocql.UUID) (*models.AlertWithPatient, error) // Método para traer alerta junto con los datos del paciente
}

type alertService struct {
	repo repository.AlertRepository
}

func NewAlertService(repo repository.AlertRepository) AlertService {
	return &alertService{repo: repo}
}

func (s *alertService) CreateAlert(alert *models.Alert) error {
	if alert.Auditable.CreatedAt.IsZero() {
		alert.Auditable.CreatedAt = time.Now()
		alert.Auditable.ModifiedAt = alert.Auditable.CreatedAt
	}
	if alert.AlertTimestamp == 0 {
		alert.AlertTimestamp = time.Now().UnixNano() / int64(time.Millisecond)
	}
	return s.repo.Create(alert)
}

func (s *alertService) GetAllAlerts() ([]*models.Alert, error) {
	return s.repo.GetAll()
}

func (s *alertService) GetAlertByID(id gocql.UUID) (*models.Alert, error) {
	return s.repo.GetByID(id)
}

func (s *alertService) UpdateAlert(alert *models.Alert) error {
	return s.repo.Update(alert)
}

func (s *alertService) DeleteAlert(id gocql.UUID) error {
	return s.repo.SoftDelete(id)
}

func (s *alertService) GetAlertWithPatient(id gocql.UUID) (*models.AlertWithPatient, error) {
	return s.repo.GetAlertWithPatient(id)
}
