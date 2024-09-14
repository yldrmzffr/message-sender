package health

import "time"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HealthCheck() (HealthResponse, error) {
	timestamp := time.Now().UTC().Format(time.RFC3339)

	return HealthResponse{
		Message:   "OK",
		Timestamp: timestamp,
	}, nil

}
