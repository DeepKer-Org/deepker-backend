package repository

import (
	"biometric-data-backend/models"
	"errors"
	"fmt"
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
	fmt.Printf("PIPIPIPIPI")
	var alert models.Alert
	if err := r.db.Preload("AttendedBy").Where(primaryKey+" = ?", id).First(&alert).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	fmt.Printf("Alerta : %v\n", alert)
	fmt.Printf("Alerta attended by: %v\n", alert.AttendedBy)
	return &alert, nil
}

func (r *alertRepository) Create(entity *models.Alert) error {
	return r.baseRepo.Create(entity)
}

func (r *alertRepository) GetAll() ([]*models.Alert, error) {
	return r.baseRepo.GetAll()
}

func (r *alertRepository) Update(entity *models.Alert, primaryKey string, id interface{}) error {
	return r.baseRepo.Update(entity, primaryKey, id)
}

func (r *alertRepository) Delete(id interface{}, primaryKey string) error {
	return r.baseRepo.Delete(id, primaryKey)
}
