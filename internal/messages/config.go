package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	redislib "github.com/redis/go-redis/v9"
	"message-sender/internal/notification"
)

func ConfigureMessagesModule(router *gin.Engine, db *pgxpool.Pool, ns notification.Provider, rdsCli *redislib.Client) {
	messagesRepo := NewRepository(db)
	messagesService := NewService(messagesRepo, ns, rdsCli)
	messagesHandler := NewMessageHandler(messagesService)

	messagesGroup := router.Group("/messages")
	RegisterRoutes(messagesGroup, messagesHandler)
}
