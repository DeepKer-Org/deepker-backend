package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type AuthorizationRepository interface {
	BaseRepository[models.User]
	GetUserByUsername(email string) (*models.User, error)
}

type authorizationRepository struct {
	BaseRepository[models.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) AuthorizationRepository {
	baseRepo := NewBaseRepository[models.User](db)
	return &authorizationRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetUserByUsername retrieves a user by their email.
func (r *authorizationRepository) GetUserByUsername(email string) (*models.User, error) {
	var user models.User
	if err := r.db.
		Preload("Roles").
		Where("username = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authorizationRepository) GetByID(id interface{}, primaryKey string) (*models.User, error) {
	var user models.User
	if err := r.db.
		Preload("Roles").
		Where(primaryKey+" = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
