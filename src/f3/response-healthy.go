package f3

// HealthyResponse is returned by the health-check request.
type HealthyResponse struct {
	// Status is the status returned by the service.
	Status *string `json:"status,omitempty"`
}
