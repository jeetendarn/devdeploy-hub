package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/jeetendar/devdeploy-hub/internal/handlers"
)

func Register(router *gin.Engine) {

	api := router.Group("/api")

	{
		api.GET("/health", handlers.Health)
		api.GET("/projects", handlers.GetProjects)
		api.POST("/projects", handlers.CreateProject)
	}
}
