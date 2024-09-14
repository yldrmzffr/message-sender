package messages

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, messageHandler *Handler) {
	router.POST("/", messageHandler.CreateMessage)
}
