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
