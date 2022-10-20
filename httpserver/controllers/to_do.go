package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/params"
	"hacktiv8-golang-final-project-1-team-8/httpserver/services"
	"net/http"
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
	var req params.ToDoCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = validator.New().Struct(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	WriteJsonResponse(ctx, t.svc.CreateToDo(&req))
}

func (t *ToDoController) UpdateToDo(ctx *gin.Context) {

}

func (t *ToDoController) DeleteToDo(ctx *gin.Context) {

}
