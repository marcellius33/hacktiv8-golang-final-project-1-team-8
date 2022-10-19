package controllers

import (
	"github.com/gin-gonic/gin"
	"hacktiv8-golang-final-project-1-team-8/httpserver/controllers/views"
)

func WriteJsonResponse(ctx *gin.Context, resp *views.Response) {
	ctx.JSON(resp.Status, resp)
}
