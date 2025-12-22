package appService

import "context"

//encore:api auth method=GET path=/health
func (s *ServiceApp) HealthCheck(ctx context.Context) (HealthResponse, error) {
	return HealthResponse{Status: "OK"}, nil
}

type HealthResponse struct {
	Status string `json:"status"`
}
