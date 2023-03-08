package entities

import (
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	Blogs     []Blog    `gorm:"foreignKey:UserID" json:"blogs,omitempty"`
	Likes     []Like    `gorm:"foreignKey:UserID" json:"likes,omitempty"`
	Comments  []Comment `gorm:"foreignKey:UserID" json:"comments,omitempty"`

	Timestamp
}
