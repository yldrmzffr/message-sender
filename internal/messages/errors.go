package messages

import (
	"message-sender/internal/pkg/apperrors"
	"net/http"
)

const (
	ErrCodeMessageDeliveryInProgress = 3001
	ErrCodeNoActiveMessageDelivery   = 3002
	ErrCodeMessageDeliveryFailed     = 3003
)

var (
	ErrMessageDeliveryInProgress = apperrors.NewAppError(http.StatusConflict, "Message delivery is currently in progress", ErrCodeMessageDeliveryInProgress, nil)
	ErrNoActiveMessageDelivery   = apperrors.NewAppError(http.StatusConflict, "No active message delivery found", ErrCodeNoActiveMessageDelivery, nil)
	ErrMessageDeliveryFailed     = apperrors.NewAppError(http.StatusInternalServerError, "Message delivery operation failed", ErrCodeMessageDeliveryFailed, nil)
)
