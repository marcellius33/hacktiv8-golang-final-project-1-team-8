package services

import (
	"database/sql"
	"fmt"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/params"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/views"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
	"strings"
	"time"

	"github.com/google/uuid"
)

type ToDoSvc struct {
	repo repositories.ToDoRepo
}

func NewToDoSvc(repo repositories.ToDoRepo) *ToDoSvc {
	return &ToDoSvc{
		repo: repo,
	}
}

func (t *ToDoSvc) GetAllToDos() *views.Response {
	toDos, err := t.repo.GetToDos()
	if err != nil {
		if err == sql.ErrNoRows {
			return views.DataNotFound(err)
		}
		return views.InternalServerError(err)
	}

	return views.SuccessFindResponse(parseModelToToDosGet(toDos), "GET_TO_DOS")
}

func (t *ToDoSvc) GetToDo(id string) *views.Response {
	toDo, err := t.repo.GetToDo(id)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessFindResponse(parseModelToToDoGet(toDo), "GET_TO_DO")
}

func (t *ToDoSvc) CreateToDo(req *params.ToDoCreateRequest) *views.Response {
	toDo := parseCreateRequestToModel(req)
	toDo.ID = uuid.New()
	now := time.Now()
	toDo.CreatedAt = &now
	toDo.UpdatedAt = &now

	err := t.repo.CreateToDo(toDo)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return views.DataConflict(err)
		}
		return views.InternalServerError(err)
	}
	return views.SuccessCreateResponse(toDo, "CREATE_TO_DO")
}

func (t *ToDoSvc) UpdateToDo(id string, req *params.ToDoUpdateRequest) *views.Response {
	toDo := parseUpdateRequestToModel(req)
	now := time.Now()
	toDo.UpdatedAt = &now
	fmt.Println(toDo.CompletedAt)
	err := t.repo.UpdateToDo(toDo)
	toDo, err = t.repo.GetToDo(id)
	if err != nil {
		return views.InternalServerError(err)
	}

	return views.SuccessUpdateResponse(toDo, "UPDATE_TO_DO")
}

func (t *ToDoSvc) DeleteToDo(id string) *views.Response {
	_, err := t.repo.GetToDo(id)
	if err != nil {
		return views.InternalServerError(err)
	}

	err = t.repo.DeleteToDo(id)
	if err != nil {
		return views.InternalServerError(err)
	}

	return views.SuccessDeleteResponse("DELETE SUCCES")
}

func parseModelToToDosGet(mod *[]models.ToDo) *[]views.ToDoGet {
	var toDos []views.ToDoGet
	for _, toDo := range *mod {
		toDos = append(toDos, views.ToDoGet{
		ID:          toDo.ID,
		Title:       toDo.Title,
		Description: toDo.Description,
		DueAt:       toDo.DueAt,
		CompletedAt: toDo.CompletedAt,
		DeletedAt:   toDo.DeletedAt,
		CreatedAt:   toDo.CreatedAt,
		UpdatedAt:   toDo.UpdatedAt,
		})
	}
	return &toDos
}


func parseModelToToDoGet(mod *models.ToDo) *views.ToDoGet {
	return &views.ToDoGet{
		ID:          mod.ID,
		Title:       mod.Title,
		Description: mod.Description,
		DueAt:       mod.DueAt,
		CompletedAt: mod.CompletedAt,
		DeletedAt:   mod.DeletedAt,
		CreatedAt:   mod.CreatedAt,
		UpdatedAt:   mod.UpdatedAt,
	}
}

func parseCreateRequestToModel(req *params.ToDoCreateRequest) *models.ToDo {
	return &models.ToDo{
		Title:       req.Title,
		Description: req.Description,
		DueAt:       req.DueAt,
	}
}

func parseUpdateRequestToModel(req *params.ToDoUpdateRequest) *models.ToDo{
	return &models.ToDo{
		Title:       req.Title,
		Description: req.Description,
		DueAt:       req.DueAt,
		CompletedAt:   req.CompletedAt,
	}
}
