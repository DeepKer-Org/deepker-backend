package repository

import (
	"biometric-data-backend/models"
	"gorm.io/gorm"
)

type RoleRepository interface {
	BaseRepository[models.Role]
	GetRolesByNames(roleNames []string) ([]*models.Role, error)
}

type roleRepository struct {
	BaseRepository[models.Role]
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	baseRepo := NewBaseRepository[models.Role](db)
	return &roleRepository{
		BaseRepository: baseRepo,
		db:             db,
	}
}

// GetRolesByNames retrieves roles by their role names.
func (r *roleRepository) GetRolesByNames(roleNames []string) ([]*models.Role, error) {
	var roles []*models.Role
	if err := r.db.Where("role_name IN (?)", roleNames).Find(&roles).Error; err != nil {
		return nil, err
	}
	return roles, nil
}
