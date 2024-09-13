package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"message-sender/internal/pkg/logger"
)

type PostgresConfig struct {
	DSN string
}

func NewPostgresDatabase(ctx context.Context, config *PostgresConfig) (*pgxpool.Pool, error) {
	logger.Info("PostgreSQL Starting...")

	poolConfig, err := pgxpool.ParseConfig(config.DSN)
	if err != nil {
		logger.Error("Error parsing PostgreSQL config: ", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		logger.Error("PostgreSQL Connection Error: ", err)
		return nil, err
	}

	logger.Info("PostgreSQL Connection success...")

	logger.Info("Pining to PostgreSQL...")
	err = pool.Ping(ctx)
	if err != nil {
		logger.Error("PostgreSQL Ping Error: ", err)
		return nil, err
	}

	logger.Info("PostgreSQL Ping successful")
	return pool, nil
}

func ClosePostgresConnection(pool *pgxpool.Pool) {
	logger.Info("Closing PostgreSQL connection...")
	pool.Close()
	logger.Info("PostgreSQL connection closed successfully.")
}
