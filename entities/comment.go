package entities

import (
	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BlogID    uuid.UUID `gorm:"type:uuid" json:"blog_id"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Text      string    `gorm:"type:varchar(255)" json:"text"`

	Timestamp
}
