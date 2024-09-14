package messages

import (
	"context"
	"message-sender/internal/notification"
	"message-sender/internal/pkg/apperrors"
	"message-sender/internal/pkg/logger"
)

type Service struct {
	messageRepository   *Repository
	notificationService notification.Provider
}

func NewService(repo *Repository, notificationService notification.Provider) *Service {
	return &Service{
		messageRepository:   repo,
		notificationService: notificationService,
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
