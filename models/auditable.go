package models

import "time"

// Auditable defines common fields for tracking creation, modification, and soft deletion
type Auditable struct {
	CreatedAt  time.Time  `json:"createdAt"`
	ModifiedAt time.Time  `json:"modifiedAt"`
	DeletedAt  *time.Time `json:"deletedAt,omitempty"`
}
