package health

// HealthResponse is a response for health check
// swagger:response HealthResponse
type HealthResponse struct {
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
}
