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

// ControlResponse represents the response body for control operations
// swagger:response ControlResponse
type ControlResponse struct {
	Message string `json:"message" example:"Automatic message sending has been stopped"`
}

// MessageDetailsResponse represents the details of a message
// swagger:response MessageDetailsResponse
type MessageDetailsResponse struct {
	*Message         `json:",inline" swagger:"object,Message"`
	ProviderResponse *RedisRecord `json:"providerResponse" swagger:"object,RedisRecord"`
}

// RedisRecord represents a record to be saved in Redis
// swagger:model RedisRecord
type RedisRecord struct {
	MessageId string `json:"messageId" example:"730e0f3e-663b-4962-bf86-b768290b7d49"`
	Provider  string `json:"provider" example:"gcp"`
}
