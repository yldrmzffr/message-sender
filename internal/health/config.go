package health

import "github.com/gin-gonic/gin"

func ConfigureHealthModules(router *gin.Engine) {
	healthService := NewService()
	healthHandler := NewHealthHandler(healthService)

	healthGroup := router.Group("/health")
	RegisterRoutes(healthGroup, healthHandler)
}
