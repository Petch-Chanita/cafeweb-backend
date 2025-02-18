package models

import (
	"time"

	"github.com/google/uuid"
)

type BaseModel struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"type:timestamp;default:current_timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:current_timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp;index" json:"deleted_at,omitempty"`
}
