package notification

import (
	"github.com/google/uuid"
	"message-sender/internal/pkg/logger"
)

type MockProvider struct {
	name string
}

func NewMockService() *MockProvider {
	return &MockProvider{
		"mock",
	}
}

func (p *MockProvider) Send(to string, content string) (*ProviderSuccessResponse, error) {
	logger.Debug("Sending message to mock provider. To: %s, Content: %s", to, content)

	var id = uuid.New()

	logger.Debug("Message sent successfully. Message ID: %s", id.String())

	return &ProviderSuccessResponse{
		Message:          "Accepted",
		MessageID:        id.String(),
		SelectedProvider: p.name,
	}, nil
}
