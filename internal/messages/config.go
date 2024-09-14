package messages

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConfigureMessagesModule(router *gin.Engine, db *pgxpool.Pool) {
	messagesRepo := NewRepository(db)
	messagesService := NewService(messagesRepo)
	messagesHandler := NewMessageHandler(messagesService)

	messagesGroup := router.Group("/messages")
	RegisterRoutes(messagesGroup, messagesHandler)
}
