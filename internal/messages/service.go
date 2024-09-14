package messages

import (
	"context"
	redislib "github.com/redis/go-redis/v9"
	"message-sender/config"
	"message-sender/internal/notification"
	"message-sender/internal/pkg/apperrors"
	"message-sender/internal/pkg/logger"
	"time"
)

type Service struct {
	config              config.MessagesConfig
	messageRepository   *Repository
	notificationService notification.Provider
	redisClient         *redislib.Client
	stopChan            chan struct{}
	isRunning           bool
}

func NewService(config config.MessagesConfig, repo *Repository, notificationService notification.Provider, redisClient *redislib.Client) *Service {
	s := &Service{
		config:              config,
		messageRepository:   repo,
		notificationService: notificationService,
		redisClient:         redisClient,
		stopChan:            make(chan struct{}),
		isRunning:           false,
	}

	if config.AutoStart == true {
		go s.StartSendingMessages(context.Background())
	}

	return s
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

func (s *Service) GetPendingMessages(ctx context.Context, limit int) ([]*Message, error) {
	logger.Debug("Getting pending messages")
	messages, err := s.messageRepository.GetPendingMessages(ctx, &limit)
	if err != nil {
		return nil, apperrors.ErrorInternalServer
	}

	return messages, nil
}

func (s *Service) ControlMessageSending(ctx context.Context, action string) error {
	logger.Debug("Control message sending", action)
	switch action {
	case "start":
		return s.StartSendingMessages(ctx)
	case "stop":
		return s.StopSendingMessages(ctx)
	default:
		return apperrors.ErrorBadRequest
	}
}

func (s *Service) StartSendingMessages(ctx context.Context) error {
	logger.Debug("Starting message delivery")
	if s.isRunning {
		return ErrMessageDeliveryInProgress
	}

	s.isRunning = true
	s.stopChan = make(chan struct{})
	go s.messageLoop(ctx)

	logger.Debug("Message delivery started")
	return nil
}

func (s *Service) StopSendingMessages(ctx context.Context) error {
	logger.Debug("Stopping message delivery")
	if !s.isRunning {
		return ErrNoActiveMessageDelivery
	}

	s.isRunning = false
	close(s.stopChan)

	logger.Debug("Message delivery stopped")
	return nil
}

func (s *Service) SendMessages(ctx context.Context) error {
	logger.Debug("Sending messages")
	messages, err := s.GetPendingMessages(ctx, 2)
	if err != nil {
		return apperrors.ErrorInternalServer
	}

	if len(messages) == 0 {
		logger.Info("No messages to send")
	}

	for _, message := range messages {
		s.sendAndUpdateMessage(ctx, message)
	}

	return nil
}

func (s *Service) sendAndUpdateMessage(ctx context.Context, message *Message) {
	logger.Debug("Sending message", message)
	response, err := s.notificationService.Send(message.Recipient, message.Content)
	if err != nil {
		logger.Error("Failed to send message", err)
		return
	}

	logger.Debug("Message sent", response)

	err = s.messageRepository.SetSentStatusAndUpdateCompletedAt(ctx, message.ID)
	if err != nil {
		logger.Error("Failed to update message status", err)
		return
	}

	err = s.SaveReceivedMessageToRedis(ctx, message, response)
	if err != nil {
		logger.Error("Failed to save received message to redis", err)
		return
	}

	return
}

func (s *Service) messageLoop(ctx context.Context) {
	internal := time.Duration(s.config.Interval) * time.Second

	ticker := time.NewTicker(internal)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if s.isRunning {
				if err := s.SendMessages(ctx); err != nil {
					logger.Error("Error sending messages", "error", err)
				}
			}
		case <-s.stopChan:
			logger.Debug("Message loop stopped")
			return
		case <-ctx.Done():
			logger.Debug("Context done, stopping message loop")
			return
		}
	}
}

func (s *Service) SaveReceivedMessageToRedis(ctx context.Context, message *Message, response *notification.ProviderSuccessResponse) error {
	logger.Debug("Saving received message to redis", message)

	const redisPrefix = "RECEIVED_MESSAGES"

	err := s.redisClient.HSet(ctx, redisPrefix, message.ID, response.MessageID).Err()
	if err != nil {
		logger.Error("Failed to save received message to redis", err)
		return apperrors.ErrorInternalServer
	}

	logger.Debug("Received message saved to redis", message)

	return nil
}
