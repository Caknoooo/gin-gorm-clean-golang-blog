package dto

import "github.com/google/uuid"

type BlogCreateDTO struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    `gorm:"type:varchar(255)" form:"title" json:"title" binding:"required"`
	Slug   string    `gorm:"type:varchar(255)" form:"slug" json:"slug" binding:"required"`
	Content string    `gorm:"type:varchar(255)" form:"content" json:"content" binding:"required"`
	UserID  uuid.UUID `gorm:"type:uuid" form:"user_id" json:"user_id" binding:"required"`
}

type BlogUpdateDTO struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Title   string    `gorm:"type:varchar(255)" form:"title" json:"title" binding:"required"`
	Slug   string    `gorm:"type:varchar(255)" form:"slug" json:"slug" binding:"required"`
	Content string    `gorm:"type:varchar(255)" form:"content" json:"content" binding:"required"`
	UserID  uuid.UUID `gorm:"type:uuid" form:"user_id" json:"user_id" binding:"required"`
}