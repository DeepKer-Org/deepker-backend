package comorbidity

import (
	"biometric-data-backend/internal/shared"
	"gorm.io/gorm"
)

type Repository interface {
	shared.BaseRepository[Model]
}

type repository struct {
	shared.BaseRepository[Model]
}

func NewRepository(db *gorm.DB) Repository {
	baseRepo := shared.NewBaseRepository[Model](db)
	return &repository{baseRepo}
}
