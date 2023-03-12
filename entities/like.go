package entities

import (
	"github.com/google/uuid"
)

type Like struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BlogID uuid.UUID `gorm:"type:uuid;not null" json:"blog_id"`
	UserID uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`

	Timestamp
}
