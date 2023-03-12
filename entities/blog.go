package entities

import (
	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title     string    `gorm:"type:varchar(255)" json:"title"`
	Slug      string    `gorm:"type:varchar(255)" json:"slug"`
	Content   string    `gorm:"type:varchar(255)" json:"content"`
	LikeCount uint64    `json:"like_count"`

	User     User      `gorm:"foreignKey:UserID" json:"user"`
	UserID   uuid.UUID `gorm:"type:uuid" json:"user_id"`
	Likes    []Like    `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"likes"`
	Comments []Comment `gorm:"foreignKey:BlogID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"comments"`

	Timestamp
}
