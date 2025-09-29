package utils

import (
	"encoding/json"
	"net/http"
)

// SuccessResponse represents a successful API reponse
type SuccessResponse struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

// ErrorResponse represents an error API reponse
type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// RespondJSON sends a JSON response
func RespondJSON(w http.ResponseWriter, statusCode int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// RespondSuccess sends a success response
func RespondSuccess(w http.ResponseWriter, statusCode int, data any, message string) {
	RespondJSON(w, statusCode, SuccessResponse{
		Success: true,
		Data:    data,
		Message: message,
	})
}

// RespondError sends an error response
func RespondError(w http.ResponseWriter, statusCode int, errorMsg string) {
	RespondJSON(w, statusCode, ErrorResponse{
		Success: false,
		Error:   errorMsg,
	})
}
