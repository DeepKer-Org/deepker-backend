package service

import (
	"biometric-data-backend/models"
	"biometric-data-backend/repository"
)

type AlertService interface {
	GetAllAlerts() ([]models.Alert, error)
	GetAlertByID(id string) (models.Alert, error)
	GetAlertByStatus(status string) ([]models.Alert, error)
	CreateAlert(alert models.Alert) (models.Alert, error)
	UpdateAlert(alert models.Alert) (models.Alert, error)
	DeleteAlert(id string) error
}

type alertService struct {
	repo repository.AlertRepository
}

func NewAlertService(repo repository.AlertRepository) AlertService {
	return &alertService{repo: repo}
}

func (s *alertService) GetAllAlerts() ([]models.Alert, error) {
	alerts, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (s *alertService) GetAlertByID(id string) (models.Alert, error) {
	alert, err := s.repo.FindByID(id)
	if err != nil {
		return models.Alert{}, err
	}
	return alert, nil
}

func (s *alertService) GetAlertByStatus(status string) ([]models.Alert, error) {
	alerts, err := s.repo.FindByStatus(status)
	if err != nil {
		return nil, err
	}
	return alerts, nil
}

func (s *alertService) CreateAlert(alert models.Alert) (models.Alert, error) {
	createdAlert, err := s.repo.Create(alert)
	if err != nil {
		return models.Alert{}, err
	}
	return createdAlert, nil
}

func (s *alertService) UpdateAlert(alert models.Alert) (models.Alert, error) {
	updatedAlert, err := s.repo.Update(alert)
	if err != nil {
		return models.Alert{}, err
	}
	return updatedAlert, nil
}

func (s *alertService) DeleteAlert(id string) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
