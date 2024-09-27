package repository

import (
	"biometric-data-backend/models"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ComorbidityRepository interface {
	CreateComorbidity(comorbidity *models.Comorbidity) error
	GetComorbidityByID(id uuid.UUID) (*models.Comorbidity, error)
	GetAllComorbidities() ([]*models.Comorbidity, error)
	UpdateComorbidity(comorbidity *models.Comorbidity) error
	DeleteComorbidity(id uuid.UUID) error
}

type comorbidityRepository struct {
	db *gorm.DB
}

func NewComorbidityRepository(db *gorm.DB) ComorbidityRepository {
	return &comorbidityRepository{db}
}

// CreateComorbidity creates a new comorbidity record in the database.
func (r *comorbidityRepository) CreateComorbidity(comorbidity *models.Comorbidity) error {
	if err := r.db.Create(comorbidity).Error; err != nil {
		return err
	}
	return nil
}

// GetComorbidityByID retrieves a comorbidity by their ComorbidityID.
func (r *comorbidityRepository) GetComorbidityByID(id uuid.UUID) (*models.Comorbidity, error) {
	var comorbidity models.Comorbidity
	if err := r.db.Where("comorbidity_id = ?", id).First(&comorbidity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comorbidity, nil
}

// GetAllComorbidities retrieves all comorbidities from the database.
func (r *comorbidityRepository) GetAllComorbidities() ([]*models.Comorbidity, error) {
	var comorbidities []*models.Comorbidity
	if err := r.db.Find(&comorbidities).Error; err != nil {
		return nil, err
	}
	return comorbidities, nil
}

// UpdateComorbidity updates an existing comorbidity record in the database.
func (r *comorbidityRepository) UpdateComorbidity(comorbidity *models.Comorbidity) error {
	if err := r.db.Save(comorbidity).Error; err != nil {
		return err
	}
	return nil
}

// DeleteComorbidity deletes a comorbidity by their ComorbidityID.
func (r *comorbidityRepository) DeleteComorbidity(id uuid.UUID) error {
	if err := r.db.Where("comorbidity_id = ?", id).Delete(&models.Comorbidity{}).Error; err != nil {
		return err
	}
	return nil
}
