package messages

import (
	"github.com/gin-gonic/gin"
	"message-sender/internal/pkg/apperrors"
)

type Handler struct {
	service *Service
}

func NewMessageHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// CreateMessage @Summary Create a new message
// @Summary Create a new message
// @Description Create a new message
// @Tags Message
// @Accept json
// @Produce json
// @Param message body CreateMessageRequest true "Message to create"
// @Success 201 {object} MessageResponse "Created"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /messages [post]
func (h *Handler) CreateMessage(ctx *gin.Context) {
	var message CreateMessageRequest
	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.Error(apperrors.ErrorValidation)
		return
	}

	createdMessage, err := h.service.CreateMessage(ctx, &message)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(201, createdMessage)
}
