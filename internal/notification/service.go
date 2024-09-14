package notification

const (
	MockSms = "mock"
	GCPSms  = "gcp"
)

type ProviderSuccessResponse struct {
	Message          string `json:"message"`
	MessageID        string `json:"messageId"`
	SelectedProvider string `json:"selectedProvider"`
}

type Provider interface {
	Send(to string, content string) (*ProviderSuccessResponse, error)
}

type Notification struct {
	provider Provider
}

func NewService(providerName string) (Provider, error) {
	var provider Provider

	switch providerName {
	case GCPSms:
		provider = NewGCPSmsProvider()
	case MockSms:
		provider = NewMockService()
	default:
		return nil, ErrProviderNotFoundResponse
	}

	return provider, nil
}
