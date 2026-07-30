package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jeetendar/devdeploy-hub/internal/models"
	"github.com/jeetendar/devdeploy-hub/internal/repository"
)

func GetProjects(c *gin.Context) {

	projects, err := repository.GetProjects()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, projects)
}

func CreateProject(c *gin.Context) {

	var project models.Project

	if err := c.ShouldBindJSON(&project); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := repository.CreateProject(project); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Project created successfully",
	})
}