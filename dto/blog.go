package dto

import "github.com/google/uuid"

type BlogDTO struct {
	ID      uuid.UUID `json:"id"`
	Blog    string    `json:"blog"`
	Like    uint64    `json:"like"`
	Comment string    `json:"comment"`
	Author  string    `json:"author"`
	UserID  uuid.UUID `json:"user_id"`
}
