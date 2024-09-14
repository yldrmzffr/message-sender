package notification

import (
	"message-sender/internal/pkg/apperrors"
	"net/http"
)

const (
	ErrProviderNotFound = 2001
)

var (
	ErrProviderNotFoundResponse = apperrors.NewAppError(http.StatusInternalServerError, "Notification provider not found", ErrProviderNotFound, nil)
)
