package redis

import (
	"context"
	redislib "github.com/redis/go-redis/v9"
	"message-sender/internal/pkg/logger"
)

type Config struct {
	Url string
}

func NewRedis(ctx context.Context, config *Config) (*redislib.Client, error) {
	logger.Info("Redis Starting...")
	redisOpt, err := redislib.ParseURL(config.Url)
	if err != nil {
		logger.Error("Error parsing Redis config: ", err)
		return nil, err
	}

	client := redislib.NewClient(redisOpt)

	logger.Info("Pinging to Redis...")
	_, err = client.Ping(ctx).Result()
	if err != nil {
		logger.Error("Redis Ping Error: ", err)
		return nil, err
	}

	logger.Info("Redis Ping successful")
	logger.Info("Redis Connection success...")

	return client, nil
}

func CloseRedisConnection(client *redislib.Client) {
	logger.Info("Closing Redis connection...")
	client.Close()
	logger.Info("Redis connection closed successfully.")
}
