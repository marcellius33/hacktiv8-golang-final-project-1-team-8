package repositories

import (
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
)

type ToDoRepo interface {
	GetToDos() (*[]models.ToDo, error)
	GetToDo(id string) (*models.ToDo, error)
	CreateToDo(toDo *models.ToDo) error
	UpdateToDo(toDo *models.ToDo) error
	DeleteToDo(id string) error
}
