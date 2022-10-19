package main

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-golang-final-project-1-team-8/config"
	"hacktiv8-golang-final-project-1-team-8/httpserver"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/postgres"
	"hacktiv8-golang-final-project-1-team-8/httpserver/services"
)

func main() {
	db, err := config.ConnectPostgresql()
	if err != nil {
		panic(err)
	}
	toDoRepo := postgres.NewToDoRepo(db)
	toDoSvc := services.NewToDoSvc(toDoRepo)
	toDoHandler := controllers.NewToDoController(toDoSvc)

	router := gin.Default()

	app := httpserver.NewRouter(router, toDoHandler)
	app.Start(":8080")
}
