package services

import (
	"github.com/google/uuid"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/params"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/views"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
	"strings"
	"time"
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
	//TODO implement me
	panic("implement me")
}

func (t *ToDoSvc) GetToDo(id string) *views.Response {
	toDo, err := t.repo.GetToDo(id)
	if err != nil {
		return views.InternalServerError(err)
	}
	return views.SuccessFindResponse(parseModelToToDoGet(toDo), "GET_TO_DO")
}

func (t *ToDoSvc) CreateToDo(req *params.ToDoCreateRequest) *views.Response {
	toDo := parseRequestToModel(req)
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

func (t *ToDoSvc) UpdateToDo() *views.Response {
	//TODO implement me
	panic("implement me")
}

func (t *ToDoSvc) DeleteToDo() *views.Response {
	//TODO implement me
	panic("implement me")
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

func parseRequestToModel(req *params.ToDoCreateRequest) *models.ToDo {
	return &models.ToDo{
		Title:       req.Title,
		Description: req.Description,
		DueAt:       req.DueAt,
	}
}
