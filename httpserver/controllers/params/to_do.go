package params

import (
	"time"
)

type ToDoCreateRequest struct {
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	DueAt       *time.Time `json:"due_at"`
}

type ToDoUpdateRequest struct {
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	DueAt       *time.Time `json:"due_at"`
	CompletedAt *time.Time `json:"completed_at"`
}