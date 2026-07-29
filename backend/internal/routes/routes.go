package routes

import (
	"devdeploy-hub/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {

	api := router.Group("/api")

	{

		api.GET("/health", handlers.Health)

		api.GET("/projects", handlers.GetProjects)

		api.POST("/projects", handlers.CreateProject)

	}

}