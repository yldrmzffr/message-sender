package health

import "github.com/gin-gonic/gin"

type Handler struct {
	service *Service
}

func NewHealthHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// HealthCheck @Summary Health check
// @Summary Health check
// @Description Health check
// @Tags Health
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse "OK"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /health [get]
func (h *Handler) HealthCheck(ctx *gin.Context) {
	health, err := h.service.HealthCheck()
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(200, health)
}
