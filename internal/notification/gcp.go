package notification

type GCPSmsProvider struct {
	name string
}

func NewGCPSmsProvider() *GCPSmsProvider {
	return &GCPSmsProvider{
		"GCPSmsProvider",
	}
}

func (p *GCPSmsProvider) Send(to string, content string) (*ProviderSuccessResponse, error) {
	// todo: implement GCP SMS sending logic

	return &ProviderSuccessResponse{
		Message:          "SMS sent successfully",
		MessageID:        "123",
		SelectedProvider: p.name,
	}, nil
}
