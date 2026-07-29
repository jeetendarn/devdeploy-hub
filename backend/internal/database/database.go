package database

import (
	"context"
	"fmt"
	"log"

"github.com/jeetendar/devdeploy-hub/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect(cfg *config.Config) error {

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	db, err := pgxpool.New(context.Background(), dsn)

	if err != nil {
		return err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return err
	}

	DB = db

	log.Println("PostgreSQL Connected")

	return nil

}