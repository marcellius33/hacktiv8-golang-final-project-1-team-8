package postgres

import (
	"database/sql"
	"fmt"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
)

type toDoRepo struct {
	db *sql.DB
}

func (t toDoRepo) GetToDos() (*[]models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t toDoRepo) GetToDo(id string) (*models.ToDo, error) {
	var toDo models.ToDo
	if err := t.db.QueryRow("SELECT * FROM to_dos WHERE id = $1", id).Scan(&toDo.ID, &toDo.Title, &toDo.Description, &toDo.DueAt, &toDo.CompletedAt, &toDo.DeletedAt, &toDo.CreatedAt, &toDo.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid id : %s", id)
		}
		return nil, err
	}
	return &toDo, nil
}

func (t toDoRepo) CreateToDo(toDo *models.ToDo) error {
	stmt, err := t.db.Prepare("INSERT INTO to_dos (id, title, description, due_at, completed_at, deleted_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(toDo.ID, toDo.Title, toDo.Description, toDo.DueAt, toDo.CompletedAt, toDo.DeletedAt, toDo.CreatedAt, toDo.UpdatedAt)

	return err
}

func (t toDoRepo) UpdateToDo(toDo *models.ToDo) (*models.ToDo, error) {
	//TODO implement me
	panic("implement me")
}

func (t toDoRepo) DeleteToDo(toDo *models.ToDo) error {
	//TODO implement me
	panic("implement me")
}

func NewToDoRepo(db *sql.DB) repositories.ToDoRepo {
	return &toDoRepo{
		db: db,
	}
}
