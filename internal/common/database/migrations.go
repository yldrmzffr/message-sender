package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"message-sender/internal/pkg/logger"
)

func RunMigrations(databaseURL string) error {
	m, err := migrate.New(
		"file://migrations",
		databaseURL)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	logger.Info("Migrations completed successfully")
	return nil
}
