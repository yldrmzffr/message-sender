package messages

import (
	"context"
	"message-sender/internal/pkg/apperrors"
	"message-sender/internal/pkg/logger"
)

type Service struct {
	messageRepository *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		messageRepository: repo,
	}
}

func (s *Service) CreateMessage(ctx context.Context, message *CreateMessageRequest) (*Message, error) {
	logger.Debug("Creating message", message)
	createdMessage, err := s.messageRepository.Create(ctx, message)
	if err != nil {
		return nil, apperrors.ErrorInternalServer
	}

	return createdMessage, nil
}

func (s *Service) GetSentMessages(ctx context.Context) ([]*Message, error) {
	logger.Debug("Getting sent messages")
	messages, err := s.messageRepository.GetSentMessages(ctx)
	if err != nil {
		return nil, apperrors.ErrorInternalServer
	}

	return messages, nil
}
