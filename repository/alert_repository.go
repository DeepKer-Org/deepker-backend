package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
	"time"
)

type AlertRepository interface {
	FindAll() ([]models.Alert, error)
	FindByID(id string) (models.Alert, error)
	FindByStatus(status string) ([]models.Alert, error)
	Create(alert models.Alert) (models.Alert, error)
	Update(alert models.Alert) (models.Alert, error)
	Delete(id string) error
}

type alertRepository struct {
	db *gorm.DB
}

func NewAlertRepository(db *gorm.DB) AlertRepository {
	return &alertRepository{db: db}
}

// Ajustar para ignorar los eliminados lógicamente
func (r *alertRepository) FindAll() ([]models.Alert, error) {
	var alerts []models.Alert
	err := r.db.Preload("Biometrics").Preload("ComputerDiagnoses").Preload("AssociatedDoctors").Where("deleted_at IS NULL").Find(&alerts).Error
	return alerts, err
}

// Ajustar para ignorar los eliminados lógicamente
func (r *alertRepository) FindByID(id string) (models.Alert, error) {
	var alert models.Alert
	err := r.db.Preload("Biometrics").Preload("ComputerDiagnoses").Preload("AssociatedDoctors").Where("alert_id = ? AND deleted_at IS NULL", id).First(&alert).Error
	return alert, err
}

func (r *alertRepository) FindByStatus(status string) ([]models.Alert, error) {
	var alerts []models.Alert
	err := r.db.Preload("Biometrics").Preload("ComputerDiagnoses").Preload("AssociatedDoctors").Where("alert_status = ? AND deleted_at IS NULL", status).Find(&alerts).Error
	return alerts, err
}

// Crear y actualizar siguen igual
func (r *alertRepository) Create(alert models.Alert) (models.Alert, error) {
	err := r.db.Create(&alert).Error
	return alert, err
}

func (r *alertRepository) Update(alert models.Alert) (models.Alert, error) {
	err := r.db.Save(&alert).Error
	return alert, err
}

// Implementar borrado lógico
func (r *alertRepository) Delete(id string) error {
	var alert models.Alert
	err := r.db.Model(&alert).Where("alert_id = ?", id).Update("deleted_at", gorm.DeletedAt{Time: time.Now(), Valid: true}).Error
	return err
}
