package repository

import "gorm.io/gorm"

type Repository[T any] interface {
	FindAll() ([]T, error)
	FindByID(id string) (T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id string) error
}

type repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) Repository[T] {
	return &repository[T]{db: db}
}

func (r *repository[T]) FindAll() ([]T, error) {
	var entities []T
	err := r.db.Find(&entities).Error
	return entities, err
}

func (r *repository[T]) FindByID(id string) (T, error) {
	var entity T
	err := r.db.First(&entity, id).Error
	return entity, err
}

func (r *repository[T]) Create(entity *T) (*T, error) {
	err := r.db.Create(&entity).Error
	return entity, err
}

func (r *repository[T]) Update(entity *T) (*T, error) {
	err := r.db.Save(&entity).Error
	return entity, err
}

func (r *repository[T]) Delete(id string) error {
	var entity T
	return r.db.Delete(&entity, id).Error
}
