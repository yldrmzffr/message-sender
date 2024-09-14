package messages

import (
	"github.com/gin-gonic/gin"
	"message-sender/internal/pkg/apperrors"
	"time"
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

	response := MessageResponse{
		ID:        createdMessage.ID,
		Recipient: createdMessage.Recipient,
		Content:   createdMessage.Content,
		Status:    createdMessage.Status,
		CreatedAt: createdMessage.CreatedAt.Format(time.RFC3339),
	}

	ctx.JSON(201, response)
}

// GetSentMessages @Summary Get sent messages
// @Summary Get sent messages
// @Description Get sent messages
// @Tags Message
// @Accept json
// @Produce json
// @Success 200 {array} MessageResponse "Sent messages"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /messages/sent [get]
func (h *Handler) GetSentMessages(ctx *gin.Context) {
	messages, err := h.service.GetSentMessages(ctx)
	if err != nil {
		ctx.Error(err)
		return
	}

	response := make([]MessageResponse, 0)
	for _, message := range messages {
		response = append(response, MessageResponse{
			ID:          message.ID,
			Recipient:   message.Recipient,
			Content:     message.Content,
			Status:      message.Status,
			CreatedAt:   message.CreatedAt.Format(time.RFC3339),
			CompletedAt: message.CompletedAt.Format(time.RFC3339),
		})
	}

	ctx.JSON(200, response)
}

// ControlMessageSending @Summary Control message sending
// @Summary Control message sending
// @Description Control message sending
// @Tags Message
// @Accept json
// @Produce json
// @Param action query string true "Action to perform" Enums(start, stop)
// @Success 200 {object} ControlResponse "Success"
// @Failure 400 {object} apperrors.ErrorResponse "Bad Request"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /messages/control [post]
func (h *Handler) ControlMessageSending(ctx *gin.Context) {
	action := ctx.Query("action")
	err := h.service.ControlMessageSending(ctx, action)
	if err != nil {
		ctx.Error(err)
		return
	}

	message := "Automatic message sending has been " + action + "ed"

	ctx.JSON(200, ControlResponse{
		Message: message,
	})
}
