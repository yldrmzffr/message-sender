package notification

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigureNotificationModule(t *testing.T) {
	tests := []struct {
		name     string
		provider string
		wantErr  bool
		expected interface{}
	}{
		{
			"GCP Provider", GCPSms, false, &GCPSmsProvider{},
		},
		{
			"Mock Provider", MockSms, false, &MockProvider{},
		},
		{
			"Invalid Provider", "invalid", true, ErrProviderNotFoundResponse,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			provider, err := ConfigureNotificationModule(test.provider)

			if test.wantErr {
				assert.Error(t, err)
				assert.Nil(t, provider)
			} else {
				assert.NoError(t, err)
				assert.IsType(t, test.expected, provider)
			}
		})
	}
}
