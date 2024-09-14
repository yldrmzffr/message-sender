package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	"message-sender/config"
	"message-sender/internal/common/database"
	"message-sender/internal/pkg/logger"
	"message-sender/internal/pkg/middleware"
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

	// Set default logger
	gin.DefaultWriter = zap.NewStdLog(logger.Log).Writer()
	gin.DefaultErrorWriter = zap.NewStdLog(logger.Log).Writer()

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

	// Gin setup
	r := gin.Default()

	r.Use(middleware.HandleError())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	port := cfg.Service.Port

	if err := r.Run(":" + port); err != nil {
		logger.Error("Error starting server", err)
		return
	}
}
