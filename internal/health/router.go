package health

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, healthHandler *Handler) {
	router.GET("/", healthHandler.HealthCheck)
}
