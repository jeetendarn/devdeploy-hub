package repository

import (
	"context"

	"devdeploy-hub/internal/database"
	"devdeploy-hub/internal/models"

	"github.com/google/uuid"
)

func GetProjects() ([]models.Project, error) {

	rows, err := database.DB.Query(
		context.Background(),
		`SELECT id,name,description,environment,status,created_at
		 FROM projects
		 ORDER BY created_at DESC`,
	)

	if err != nil {

		return nil, err

	}

	defer rows.Close()

	var projects []models.Project

	for rows.Next() {

		var p models.Project

		rows.Scan(

			&p.ID,
			&p.Name,
			&p.Description,
			&p.Environment,
			&p.Status,
			&p.CreatedAt,
		)

		projects = append(projects, p)

	}

	return projects, nil

}

func CreateProject(project models.Project) error {

	id := uuid.New()

	_, err := database.DB.Exec(

		context.Background(),

		`INSERT INTO projects
		(id,name,description,environment,status)
		VALUES($1,$2,$3,$4,$5)`,

		id,

		project.Name,

		project.Description,

		project.Environment,

		project.Status,
	)

	return err

}