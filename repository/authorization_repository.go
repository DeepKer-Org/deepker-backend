package repository

import (
	"biometric-data-backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuthorizationRepository interface {
	BaseRepository[models.User]
	GetUserByUsername(email string) (*models.User, error)
	GetUserByID(id uuid.UUID) (*models.User, error)
	GetAllUsers(offset int, limit int) ([]*models.User, error)
	CountAllUsers() (int64, error)
	DeleteUserAndUserRoles(id uuid.UUID) error
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

func (r *authorizationRepository) GetUserByID(id uuid.UUID) (*models.User, error) {
	return r.GetByID(id, "user_id")
}

func (r *authorizationRepository) GetAllUsers(offset int, limit int) ([]*models.User, error) {
	var users []*models.User

	// Chain Offset and Limit with the Preload and Find calls
	if err := r.db.
		Preload("Roles").
		Offset(offset).
		Limit(limit).
		Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *authorizationRepository) CountAllUsers() (int64, error) {
	var count int64
	if err := r.db.Model(&models.User{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (r *authorizationRepository) DeleteUserAndUserRoles(id uuid.UUID) error {
	tx := r.db.Begin()

	// Delete user roles
	if err := tx.Where("user_id = ?", id).Delete(&models.UserRole{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Delete user
	if err := tx.Where("user_id = ?", id).Delete(&models.User{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
