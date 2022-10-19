package main

import "hacktiv8-golang-final-project-1-team-8/httpserver"

func main() {
	app := httpserver.CreateRouter()
	app.Run(":8080")
}
