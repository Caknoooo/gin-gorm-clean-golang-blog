package dto

import "github.com/google/uuid"

type UserCreateDTO struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama     string    `gorm:"type:varchar(100)" json:"nama" binding:"required"`
	NoTelp   string    `gorm:"type:varchar(20)" json:"no_telp" binding:"required"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex" json:"email" binding:"required"`
	Password string    `gorm:"type:varchar(100);not null" json:"password" binding:"required"`
}

type UserUpdateDTO struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Nama     string    `gorm:"type:varchar(100)" json:"nama"`
	NoTelp   string    `gorm:"type:varchar(20)" json:"no_telp"`
	Email    string    `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password string    `gorm:"type:varchar(100);not null" json:"password"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"email"`
	Password string `json:"password" binding:"password"`
}
