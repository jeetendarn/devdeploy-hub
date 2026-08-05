package models

import "time"

type Project struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	Environment string `json:"environment"`

	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
}
