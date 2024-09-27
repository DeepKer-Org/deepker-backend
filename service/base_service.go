package service

import "biometric-data-backend/repository"

type Service[T any] interface {
	GetAll() ([]T, error)
	GetByID(id string) (T, error)
	Create(entity *T) (*T, error)
	Update(entity *T) (*T, error)
	Delete(id string) error
}

type service[T any] struct {
	repository repository.Repository[T]
}

func NewService[T any](repository repository.Repository[T]) Service[T] {
	return &service[T]{repository: repository}
}

func (s *service[T]) GetAll() ([]T, error) {
	return s.repository.FindAll()
}

func (s *service[T]) GetByID(id string) (T, error) {
	return s.repository.FindByID(id)
}

func (s *service[T]) Create(entity *T) (*T, error) {
	return s.repository.Create(entity)
}

func (s *service[T]) Update(entity *T) (*T, error) {
	return s.repository.Update(entity)
}

func (s *service[T]) Delete(id string) error {
	return s.repository.Delete(id)
}
