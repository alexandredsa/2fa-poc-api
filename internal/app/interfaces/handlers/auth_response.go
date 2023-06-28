package handlers

import "github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"

// RegisterResponse represents the response body for user registration.
type RegisterResponse struct {
	User    models.User `json:"user"`
	Message string      `json:"message"`
}
