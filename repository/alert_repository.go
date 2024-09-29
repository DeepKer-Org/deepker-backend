package repository

import (
	"biometric-data-backend/models"
	"errors"
	"gorm.io/gorm"
)

// AlertRepository define los métodos específicos para la entidad Alert
type AlertRepository interface {
	BaseRepository[models.Alert] // Incluir métodos del repositorio base
}

// alertRepository implementa AlertRepository
type alertRepository struct {
	baseRepo BaseRepository[models.Alert] // Delegamos los métodos al baseRepository
	db       *gorm.DB
}

// NewAlertRepository crea una nueva instancia de AlertRepository
func NewAlertRepository(db *gorm.DB) AlertRepository {
	return &alertRepository{
		baseRepo: NewBaseRepository[models.Alert](db), // Inicializamos el baseRepo
		db:       db,
	}
}

func (r *alertRepository) GetByID(id interface{}, primaryKey string) (*models.Alert, error) {
	var alert models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Where(primaryKey+" = ?", id).First(&alert).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &alert, nil
}

func (r *alertRepository) Create(entity *models.Alert) error {
	return r.baseRepo.Create(entity)
}

func (r *alertRepository) GetAll() ([]*models.Alert, error) {
	var alerts []*models.Alert
	if err := r.db.Preload("BiometricData").
		Preload("AttendedBy").
		Preload("Patient").
		Find(&alerts).Error; err != nil {
		return nil, err
	}
	return alerts, nil
}

func (r *alertRepository) Update(entity *models.Alert, primaryKey string, id interface{}) error {
	return r.baseRepo.Update(entity, primaryKey, id)
}

func (r *alertRepository) Delete(id interface{}, primaryKey string) error {
	return r.baseRepo.Delete(id, primaryKey)
}
