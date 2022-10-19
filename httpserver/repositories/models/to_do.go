package models

import (
	"github.com/google/uuid"
	"time"
)

type ToDo struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CompletedAt time.Time `json:"completed_at"`
	DeletedAt   time.Time `json:"deleted_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
