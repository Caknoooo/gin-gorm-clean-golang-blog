package entities

import (
	"github.com/google/uuid"
)

type Blog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Blog      string    `gorm:"type:varchar(255)" json:"blog"`
	Like      uint64    `gorm:"type:uint64" json:"like"`
	Comment   string    `gorm:"type:varchar(255)" json:"comment"`
	Author    string    `gorm:"type:varchar(255)" json:"author"`
	UserID    uuid.UUID `gorm:"type:uuid" json:"user_id"`
	User      User      `gorm:"foreignKey:UserID" json:"user"`
	Likes     []Like    `gorm:"foreignKey:BlogID" json:"likes"`
	Comments  []Comment `gorm:"foreignKey:BlogID" json:"comments"`

	Timestamp
}
