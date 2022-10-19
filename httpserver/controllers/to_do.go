package controllers

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-golang-final-project-1-team-8/httpserver/services"
)

type ToDoController struct {
	svc *services.ToDoSvc
}

func NewToDoController(svc *services.ToDoSvc) *ToDoController {
	return &ToDoController{
		svc: svc,
	}
}

func (t *ToDoController) GetAllToDos(ctx *gin.Context) {

}

func (t *ToDoController) GetToDo(ctx *gin.Context) {
	WriteJsonResponse(ctx, t.svc.GetToDo(ctx.Param("id")))
}

func (t *ToDoController) CreateToDo(ctx *gin.Context) {

}

func (t *ToDoController) UpdateToDo(ctx *gin.Context) {

}

func (t *ToDoController) DeleteToDo(ctx *gin.Context) {

}
