package main

import (
	"context"
	"github.com/kelseyhightower/envconfig"
	"message-sender/config"
	"message-sender/internal/common/database"
	"message-sender/internal/pkg/logger"
)

var (
	cfg config.Config
)

func configSetup() {
	err := envconfig.Process("", &cfg)
	if err != nil {
		logger.Error("Config setup error", err)
		return
	}
}

func migrateDatabase(cfg *config.DatabaseConfig) {
	dbUrl := cfg.GetURL()
	if err := database.RunMigrations(dbUrl); err != nil {
		logger.Error("Database migrations error", err)
		return
	}
}

func main() {
	ctx := context.Background()

	configSetup()

	logger.InitLogger(&cfg.Log)

	logger.Info("Starting application...")

	// Migrate database
	migrateDatabase(&cfg.Database)

	// Database connection
	db, err := database.NewPostgresDatabase(ctx, &database.PostgresConfig{DSN: cfg.Database.GetDSN()})
	if err != nil {
		logger.Error("Database Connection Error", err)
		return
	}

	defer database.ClosePostgresConnection(db)

}
