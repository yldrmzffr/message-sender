package notification

import (
	"message-sender/internal/pkg/apperrors"
	"net/http"
)

const (
	ErrProviderNotFound = 2001
	ErrorSendingMessage = 2002
)

var (
	ErrProviderNotFoundResponse = apperrors.NewAppError(http.StatusInternalServerError, "Notification provider not found", ErrProviderNotFound, nil)
	ErrSendingMessageResponse   = apperrors.NewAppError(http.StatusInternalServerError, "Error while sending message", ErrorSendingMessage, nil)
)
