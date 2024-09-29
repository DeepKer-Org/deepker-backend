package dto

// CreateDTO is a generic interface for DTOs used in the create operation.
type CreateDTO[T any] interface {
	MapToEntity() *T
}

// UpdateDTO is a generic interface for DTOs used in the update operation.
type UpdateDTO[T any] interface {
	MapToEntity(entity *T) *T
}

// Mapper is a generic interface for mapping entities to DTOs.
type Mapper[T any, DTO any] interface {
	MapToDTO(entity *T) *DTO
	MapToDTOs(entities []*T) []*DTO
}
