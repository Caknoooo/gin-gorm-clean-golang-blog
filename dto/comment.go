package dto

import "github.com/google/uuid"

type CommentCreateDTO struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BlogID uuid.UUID `gorm:"type:uuid" form:"blog_id" json:"blog_id" binding:"required"`
	UserID uuid.UUID `gorm:"type:uuid" form:"user_id" json:"user_id" binding:"required"`
	Text   string    `gorm:"type:varchar(255)" form:"text" json:"text" binding:"required"`
}

type CommentUpdateDTO struct {
	ID     uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	BlogID uuid.UUID `gorm:"type:uuid" form:"blog_id" json:"blog_id" binding:"required"`
	UserID uuid.UUID `gorm:"type:uuid" form:"user_id" json:"user_id" binding:"required"`
	Text   string    `gorm:"type:varchar(255)" form:"text" json:"text" binding:"required"`
}