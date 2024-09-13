package apperrors

import (
	"fmt"
)

type AppError struct {
	Code      int
	ErrorCode int    `json:"errorCode" example:"1001"`
	Message   string `json:"message" example:"Error message"`
	Err       error  `json:"error"`
}

func (e AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(code int, message string, errorCode int, err error) *AppError {
	return &AppError{
		Code:      code,
		ErrorCode: errorCode,
		Message:   message,
		Err:       err,
	}
}
