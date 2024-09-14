package notification

func ConfigureNotificationModule(providerName string) (Provider, error) {
	provider, err := NewService(providerName)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
