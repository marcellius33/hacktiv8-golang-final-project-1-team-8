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

func NewToDoRepo(db *sql.DB) repositories.ToDoRepo {
	return &toDoRepo{
		db: db,
	}
}

func (t toDoRepo) GetToDos() (*[]models.ToDo, error) {
	rows, err := t.db.Query("SELECT * FROM to_dos")
	if err != nil {
		return &[]models.ToDo{}, err
	}
	defer rows.Close()

	var toDos []models.ToDo
	for rows.Next(){
		var toDo models.ToDo
		rows.Scan(
			&toDo.ID,
			&toDo.Title,
			&toDo.Description,
			&toDo.DueAt,
			&toDo.CompletedAt,
			&toDo.DeletedAt,
			&toDo.CreatedAt,
			&toDo.UpdatedAt,
		)

		toDos =append(toDos, toDo)
	}

	return &toDos, nil
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

func (t toDoRepo) UpdateToDo(toDo *models.ToDo)  error {
	stmt, err := t.db.Prepare("UPDATE to_dos SET title = $1, description = $2, due_at = $3, completed_at = $4, updated_at =$5 RETURNING id")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(toDo.Title, toDo.Description, toDo.DueAt, toDo.CompletedAt, toDo.UpdatedAt)

	return err
}

func (t toDoRepo) DeleteToDo(id string) error {
	_, err := t.db.Exec("DELETE FROM to_dos WHERE id = $1", id)

	return err
}
