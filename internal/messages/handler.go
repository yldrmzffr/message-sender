package messages

import (
	"github.com/gin-gonic/gin"
	"message-sender/internal/pkg/apperrors"
	"strconv"
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

	ctx.JSON(200, messages)
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

// GetMessageByID @Summary Get message by ID
// @Summary Get message by ID
// @Description Get message details by ID
// @Tags Message
// @Accept json
// @Produce json
// @Param id path int true "Message ID"
// @Success 200 {object} MessageDetailsResponse "Message details"
// @Failure 404 {object} apperrors.ErrorResponse "Not Found"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /messages/{id} [get]
func (h *Handler) GetMessageByID(ctx *gin.Context) {
	messageID := ctx.Param("id")
	id, err := strconv.Atoi(messageID)
	if err != nil {
		ctx.Error(apperrors.ErrorValidation)
		return
	}

	message, err := h.service.GetMessageByID(ctx, uint(id))
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, message)
}

// GetMessagesWithStatus @Summary Get messages with status
// @Summary Get messages with status
// @Description Get messages with status
// @Tags Message,teST
// @Accept json
// @Produce json
// @Param status query string false "Message status" Enums(sent, pending)
// @Success 200 {array} MessageResponse "Messages"
// @Failure 500 {object} apperrors.ErrorResponse "Internal Server Error"
// @Router /messages [get]
func (h *Handler) GetMessagesWithStatus(ctx *gin.Context) {
	status := ctx.Query("status")
	var messages []*Message
	var err error
	switch status {
	case string(StatusSend):
		messages, err = h.service.GetSentMessages(ctx)
	case string(StatusPending):
		messages, err = h.service.GetPendingMessages(ctx)
	default:
		ctx.Error(apperrors.ErrorValidation)
		return
	}

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(200, messages)
}
