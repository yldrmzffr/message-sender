package apperrors

import "net/http"

const (
	InternalServerErrorCode = -1
	UnauthorizedErrorCode   = -3
	BadRequestErrorCode     = -3
	ValidationErrorCode     = -4
)

var (
	ErrorInternalServer = NewAppError(http.StatusInternalServerError, "Internal server error", InternalServerErrorCode, nil)
	ErrorUnauthorized   = NewAppError(http.StatusUnauthorized, "Unauthorized", UnauthorizedErrorCode, nil)
	ErrorBadRequest     = NewAppError(http.StatusBadRequest, "Bad request", BadRequestErrorCode, nil)
	ErrorValidation     = NewAppError(http.StatusUnprocessableEntity, "Validation error", ValidationErrorCode, nil)
)
