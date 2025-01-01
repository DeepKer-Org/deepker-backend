package shared

import (
	"errors"
	"gorm.io/gorm"
)

// BaseRepository is an interface that contains common methods for all repositories
type BaseRepository[T any] interface {
	Create(entity *T) error
	CreateInTransaction(entity *T, tx *gorm.DB) error
	GetByID(id interface{}, primaryKey string) (*T, error)
	GetByIDs(ids []interface{}, primaryKey string) ([]*T, error)
	GetAll() ([]*T, error)
	Update(entity *T, primaryKey string, id interface{}) error
	UpdateInTransaction(entity *T, primaryKey string, id interface{}, tx *gorm.DB) error
	Delete(id interface{}, primaryKey string) error
	DeleteInTransaction(id interface{}, primaryKey string, tx *gorm.DB) error
	BeginTransaction() *gorm.DB
}

type baseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) BaseRepository[T] {
	return &baseRepository[T]{db}
}

// BeginTransaction starts a new transaction
func (r *baseRepository[T]) BeginTransaction() *gorm.DB {
	return r.db.Begin()
}

// Create a new record without transaction
func (r *baseRepository[T]) Create(entity *T) error {
	return r.CreateInTransaction(entity, r.db) // Reuse transaction method
}

// CreateInTransaction a new record inside a transaction
func (r *baseRepository[T]) CreateInTransaction(entity *T, tx *gorm.DB) error {
	if err := tx.Create(entity).Error; err != nil {
		return err
	}
	return nil
}

// GetByID retrieves a record by ID (ignoring logically deleted data)
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

// GetByIDs retrieves data by a list of IDs (ignoring logically deleted data)
func (r *baseRepository[T]) GetByIDs(ids []interface{}, primaryKey string) ([]*T, error) {
	var entities []*T
	if err := r.db.Where(primaryKey+" IN (?)", ids).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// GetAll retrieves all data (ignoring logically deleted data)
func (r *baseRepository[T]) GetAll() ([]*T, error) {
	var entities []*T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Update an existing record without transaction
func (r *baseRepository[T]) Update(entity *T, primaryKey string, id interface{}) error {
	return r.UpdateInTransaction(entity, primaryKey, id, r.db) // Reuse transaction method
}

// UpdateInTransaction an existing record inside a transaction
func (r *baseRepository[T]) UpdateInTransaction(entity *T, primaryKey string, id interface{}, tx *gorm.DB) error {
	if err := tx.Model(entity).Where(primaryKey+" = ?", id).Updates(entity).Error; err != nil {
		return err
	}
	return nil
}

// Delete a record by its primary key without transaction
func (r *baseRepository[T]) Delete(id interface{}, primaryKey string) error {
	return r.DeleteInTransaction(id, primaryKey, r.db) // Reuse transaction method
}

// DeleteInTransaction a record by its primary key inside a transaction
func (r *baseRepository[T]) DeleteInTransaction(id interface{}, primaryKey string, tx *gorm.DB) error {
	entity, err := r.GetByID(id, primaryKey)
	if err != nil {
		return err
	}
	if entity == nil {
		return gorm.ErrRecordNotFound
	}
	if err := tx.Where(primaryKey+" = ?", id).Delete(new(T)).Error; err != nil {
		return err
	}
	return nil
}
