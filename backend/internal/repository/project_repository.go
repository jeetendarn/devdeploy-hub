package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/jeetendar/devdeploy-hub/internal/database"
	"github.com/jeetendar/devdeploy-hub/internal/models"
)

func GetProjects() ([]models.Project, error) {

	if database.DB == nil {
		return []models.Project{}, nil
	}

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

	projects := []models.Project{}

	for rows.Next() {

		var p models.Project

		err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Environment,
			&p.Status,
			&p.CreatedAt,
		)

		if err != nil {
			return nil, err
		}

		projects = append(projects, p)
	}

	return projects, nil
}

func CreateProject(project models.Project) error {

	if database.DB == nil {
		return fmt.Errorf("database connection is not available")
	}

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
