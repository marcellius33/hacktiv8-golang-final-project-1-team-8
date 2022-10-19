package httpserver

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers"
)

type Router struct {
	router *gin.Engine
	toDo   *controllers.ToDoController
}

func NewRouter(router *gin.Engine, toDo *controllers.ToDoController) *Router {
	return &Router{
		router: router,
		toDo:   toDo,
	}
}

func (r *Router) Start(port string) {
	r.router.GET("/todos", r.toDo.GetAllToDos)
	r.router.GET("/todos/:id", r.toDo.GetToDo)
	r.router.POST("/todos", r.toDo.CreateToDo)
	r.router.PUT("/todos/:id", r.toDo.UpdateToDo)
	r.router.DELETE("/todos/:id", r.toDo.DeleteToDo)

	r.router.Run(port)
}
