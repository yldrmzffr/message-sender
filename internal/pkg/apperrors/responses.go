package apperrors

// ErrorResponse represents an API error response
// swagger:response ErrorResponse
type ErrorResponse struct {
	ErrorCode int    `json:"errorCode"`
	Message   string `json:"message"`
}

func (e AppError) ToResponse() ErrorResponse {
	return ErrorResponse{
		ErrorCode: e.ErrorCode,
		Message:   e.Message,
	}
}
