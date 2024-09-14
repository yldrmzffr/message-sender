package messages

// CreateMessageRequest represents a request to create a new message
// swagger:model CreateMessageRequest
type CreateMessageRequest struct {
	Recipient string `json:"recipient" example:"+905380000000" binding:"required,min=10,max=15"`
	Content   string `json:"content" example:"Hello, World!" binding:"required,min=5,max=140"`
}

// MessageResponse represents the response body for message operations
// swagger:response MessageResponse
type MessageResponse struct {
	ID          uint   `json:"id" example:"1"`
	Recipient   string `json:"recipient" example:"+905380000000"`
	Content     string `json:"content" example:"Hello, World!"`
	Status      Status `json:"status" example:"sent"`
	CreatedAt   string `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	CompletedAt string `json:"completedAt,omitempty" example:"2021-01-01T00:00:00Z"`
}
