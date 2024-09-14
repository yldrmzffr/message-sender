package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	redislib "github.com/redis/go-redis/v9"
	"message-sender/config"
	"message-sender/internal/notification"
)

func ConfigureMessagesModule(router *gin.Engine, cfg *config.MessagesConfig, db *pgxpool.Pool, ns notification.Provider, rdsCli *redislib.Client) {
	messagesRepo := NewRepository(db)
	messagesService := NewService(*cfg, messagesRepo, ns, rdsCli)
	messagesHandler := NewMessageHandler(messagesService)

	messagesGroup := router.Group("/messages")
	RegisterRoutes(messagesGroup, messagesHandler)
}
