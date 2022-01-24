package f3

// ErrorResponse is the response the account API sends back, when an error encountered.
type ErrorResponse struct {
	ErrorMessage string `json:"error_message"`
}
