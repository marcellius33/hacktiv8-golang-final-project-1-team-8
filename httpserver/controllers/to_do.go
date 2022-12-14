package controllers

import (
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/params"
	_ "hacktiv8-golang-final-project-1-team-8/httpserver/controllers/views"
	_ "hacktiv8-golang-final-project-1-team-8/httpserver/repositories/models"
	"hacktiv8-golang-final-project-1-team-8/httpserver/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ToDoController struct {
	svc *services.ToDoSvc
}

func NewToDoController(svc *services.ToDoSvc) *ToDoController {
	return &ToDoController{
		svc: svc,
	}
}

// GetAllToDos godoc
// @Summary      Get todos
// @Description  Get todos
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos [get]
func (t *ToDoController) GetAllToDos(ctx *gin.Context) {
	WriteJsonResponse(ctx, t.svc.GetAllToDos())
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

// UpdateToDo godoc
// @Summary Update todo
// @Description Update todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param request body params.ToDoCreateRequest true "query params"
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos/{id} [put]
func (t *ToDoController) UpdateToDo(ctx *gin.Context) {
	var req params.ToDoUpdateRequest
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

	WriteJsonResponse(ctx, t.svc.UpdateToDo(ctx.Param("id"), &req))

}

// DeleteToDo godoc
// @Summary Delete todo
// @Description Delete todo
// @Tags         todos
// @Accept       json
// @Produce      json
// @Success      200  {object}  models.ToDo
// @Failure      400  {object}  views.Response
// @Failure      404  {object}  views.Response
// @Failure      500  {object}  views.Response
// @Router       /todos/{id} [delete]
func (t *ToDoController) DeleteToDo(ctx *gin.Context) {
	WriteJsonResponse(ctx, t.svc.DeleteToDo(ctx.Param("id")))
}
