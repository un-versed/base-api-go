package models

type ApiError struct {
	Message string `json:"error_message"`
}

func NewApiError(message string) *ApiError {
	return &ApiError{Message: message}
}
