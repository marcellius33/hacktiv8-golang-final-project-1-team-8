package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/params"
	_ "hacktiv8-golang-final-project-1-team-8/httpserver/controllers/views"
	_ "hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
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

// GetToDo godoc
// @Summary Get a specific todo
// @Description Get a todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos/{id} [get]
func (t *ToDoController) GetToDo(ctx *gin.Context) {
	WriteJsonResponse(ctx, t.svc.GetToDo(ctx.Param("id")))
}

// CreateToDo godoc
// @Summary      Create new todo
// @Description  create a new task
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param request body params.ToDoCreateRequest true "query params"
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos [post]
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
