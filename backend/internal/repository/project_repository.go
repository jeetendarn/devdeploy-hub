package repository
import "fmt"
import (
	"context"

	"github.com/jeetendar/devdeploy-hub/internal/database"
	"github.com/jeetendar/devdeploy-hub/internal/models"

	"github.com/google/uuid"
)

func GetProjects() ([]models.Project, error) {

	func GetProjects() ([]models.Project, error) {

    if database.DB == nil {
        return []models.Project{}, nil
    }

    rows, err := database.DB.Query(
        context.Background(),
        "SELECT id, name, description FROM projects",
    )

    ...
}

	rows, err := database.DB.Query(
		context.Background(),
		`SELECT id,name,description,environment,status,created_at
		 FROM projects
		 ORDER BY created_at DESC`,
	)
	if database.DB == nil {
    return nil, fmt.Errorf("database connection is not available")
}

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
