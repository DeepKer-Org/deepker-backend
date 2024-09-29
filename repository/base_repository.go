package repository

import (
	"errors"
	"gorm.io/gorm"
)

type BaseRepository[T any] interface {
	Create(entity *T) error
	GetByID(id interface{}, primaryKey string) (*T, error)
	GetAll() ([]*T, error)
	Update(entity *T, primaryKey string, id interface{}) error
	Delete(id interface{}, primaryKey string) error
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db}
}

// Create a new record
func (r *baseRepository[T]) Create(entity *T) error {
	if err := r.db.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// GetByID ignoring logically deleted records (default behavior)
func (r *baseRepository[T]) GetByID(id interface{}, primaryKey string) (*T, error) {
	var entity T
	if err := r.db.Where(primaryKey+" = ?", id).First(&entity).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &entity, nil
}

// GetAll ignoring logically deleted records (default behavior)
func (r *baseRepository[T]) GetAll() ([]*T, error) {
	var entities []*T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Update an existing record
func (r *baseRepository[T]) Update(entity *T, primaryKey string, id interface{}) error {
	if err := r.db.Model(entity).Where(primaryKey+" = ?", id).Updates(entity).Error; err != nil {
		return err
	}
	return nil
}

// Delete a record by its primary key, with a variable primary key name
func (r *baseRepository[T]) Delete(id interface{}, primaryKey string) error {
	entity, err := r.GetByID(id, primaryKey)
	if err != nil {
		return err
	}
	if entity == nil {
		return gorm.ErrRecordNotFound
	}
	if err := r.db.Where(primaryKey+" = ?", id).Delete(new(T)).Error; err != nil {
		return err
	}
	return nil
}
