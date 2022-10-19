package views

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type ToDoGet struct {
	ID          uuid.UUID    `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	DueAt       sql.NullTime `json:"due_at"`
	CompletedAt sql.NullTime `json:"completed_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
}
