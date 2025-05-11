package errs

import (
	"net/http"
)

type AppError struct {
	Message string
	Code    int `json:",omitempty"`
}

func (err AppError) AsMessage() *AppError {
	return &AppError{Message: err.Message}
}
func NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}

func NewValidationError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusUnprocessableEntity}
}
