package main

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-golang-final-project-1-team-8/config"
	_ "hacktiv8-golang-final-project-1-team-8/docs"
	"hacktiv8-golang-final-project-1-team-8/httpserver"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers"
	"hacktiv8-golang-final-project-1-team-8/httpserver/repositories/postgres"
	"hacktiv8-golang-final-project-1-team-8/httpserver/services"
)

// @title           Todo Application
// @version         1.0
// @description     This is a todo list management application.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @host      localhost:8080
// @BasePath  /
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
