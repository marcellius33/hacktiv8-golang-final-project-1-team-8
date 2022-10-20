package params

import "time"

type ToDoCreateRequest struct {
	Title       string     `validate:"required" json:"title"`
	Description string     `validate:"required" json:"description"`
	DueAt       *time.Time `json:"due_at"`
}
