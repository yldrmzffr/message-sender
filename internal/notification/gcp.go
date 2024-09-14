package notification

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const URL = "https://us-central1-auto-yildirim.cloudfunctions.net/webhook"

type GCPSmsProvider struct {
	name string
}

func NewGCPSmsProvider() *GCPSmsProvider {
	return &GCPSmsProvider{
		"GCPSmsProvider",
	}
}

func (p *GCPSmsProvider) Send(to string, content string) (*ProviderSuccessResponse, error) {
	payload := map[string]string{
		"to":      to,
		"content": content,
	}

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, ErrSendingMessageResponse
	}

	contentType := "application/json"

	req, err := http.Post(URL, contentType, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, ErrSendingMessageResponse
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusAccepted {
		return nil, ErrSendingMessageResponse
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, ErrSendingMessageResponse
	}

	var response struct {
		Message   string `json:"message"`
		MessageID string `json:"messageId"`
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, ErrSendingMessageResponse
	}

	return &ProviderSuccessResponse{
		Message:          response.Message,
		MessageID:        response.MessageID,
		SelectedProvider: p.name,
	}, nil
}
