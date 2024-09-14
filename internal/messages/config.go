package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"message-sender/internal/notification"
)

func ConfigureMessagesModule(router *gin.Engine, db *pgxpool.Pool, ns notification.Provider) {
	messagesRepo := NewRepository(db)
	messagesService := NewService(messagesRepo, ns)
	messagesHandler := NewMessageHandler(messagesService)

	messagesGroup := router.Group("/messages")
	RegisterRoutes(messagesGroup, messagesHandler)
}
