package domain

type ErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error   string        `json:"error"`             // code
	Details []ErrorDetail `json:"details,omitempty"` // phrase
}
