package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	redislib "github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
	"message-sender/config"
	_ "message-sender/docs"
	"message-sender/internal/common/database"
	"message-sender/internal/common/redis"
	"message-sender/internal/messages"
	"message-sender/internal/notification"
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

func loadModules(r *gin.Engine, cfg *config.Config, db *pgxpool.Pool, rdsCli *redislib.Client) {
	ns, err := notification.ConfigureNotificationModule(cfg.Notification.Provider)
	if err != nil {
		logger.Error("Notification Provider Error", err)
		return
	}

	messages.ConfigureMessagesModule(r, &cfg.Messages, db, ns, rdsCli)
}

// @title           Message Sender API
// @version         0.1
// @description     This is a message sender service. You can use this API to send messages to users.
// @host      localhost:8080
// @BasePath  /
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
	db, err := database.NewPostgresDatabase(ctx, &database.PostgreConfig{DSN: cfg.Database.GetDSN()})
	if err != nil {
		logger.Error("Database Connection Error", err)
		return
	}

	defer database.ClosePostgresConnection(db)

	// Redis connection
	rdsCli, err := redis.NewRedis(ctx, &redis.Config{Url: cfg.Redis.Url})
	if err != nil {
		logger.Error("Redis Connection Error", err)
		return
	}

	defer redis.CloseRedisConnection(rdsCli)

	// Gin setup
	r := gin.Default()

	r.Use(middleware.HandleError())

	// Swagger setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/swagger/index.html")
	})

	// Register modules
	loadModules(r, &cfg, db, rdsCli)

	port := cfg.Service.Port

	if err := r.Run(":" + port); err != nil {
		logger.Error("Error starting server", err)
		return
	}
}
