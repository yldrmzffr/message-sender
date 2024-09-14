package messages

import "time"

type Status string

const (
	StatusPending Status = "pending"
	StatusSend    Status = "sent"
	StatusFailed  Status = "failed"
)

// Message represents a message in the system
// @Description Message information
// swagger:model Message
type Message struct {
	ID          uint       `json:"id" example:"1"`
	Recipient   string     `json:"recipient" example:"+905380000000"`
	Content     string     `json:"content" example:"Hello, World!"`
	Status      Status     `json:"status" example:"sent"`
	CreatedAt   time.Time  `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	CompletedAt *time.Time `json:"completedAt,omitempty" example:"2021-01-01T00:00:00Z"`
}
