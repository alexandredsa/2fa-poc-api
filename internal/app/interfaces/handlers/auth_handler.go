package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
)

// AuthHandler handles the authentication-related HTTP requests.
type AuthHandler struct {
	authService      *services.AuthenticationService
	componentService *services.ComponentService
	twoFADataHandler *TwoFADataHandler
}

// NewAuthHandler creates a new instance of the AuthHandler.
func NewAuthHandler(authService *services.AuthenticationService, componentService *services.ComponentService) *AuthHandler {
	return &AuthHandler{
		authService:      authService,
		componentService: componentService,
		twoFADataHandler: NewTwoFADataHandler(componentService),
	}
}

// Register handles the registration request.
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create a new user based on the request data
	user := &models.User{
		Username: request.Username,
		Password: request.Password,
	}

	// Call the AuthService to create the user
	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	// Prepare the response
	response := RegisterResponse{
		User:    *createdUser,
		Message: "Registration successful",
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the response JSON to the response writer
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// Login handles the login request.
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Implement login logic
}

// RequestTwoFA handles the request for 2FA code.
func (h *AuthHandler) RequestTwoFA(w http.ResponseWriter, r *http.Request) {
	// Implement 2FA request logic
}

// ValidateTwoFA handles the validation of 2FA code.
func (h *AuthHandler) ValidateTwoFA(w http.ResponseWriter, r *http.Request) {
	// Implement 2FA validation logic
}

// UpdateCredentials handles the update of user credentials.
func (h *AuthHandler) UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	// Implement update credentials logic
}

// UpdateComponentData handles the update of component data.
func (h *AuthHandler) UpdateComponentData(w http.ResponseWriter, r *http.Request) {
	// Implement update component data logic
}

// ValidateComponentData handles the validation of component data.
func (h *AuthHandler) ValidateComponentData(w http.ResponseWriter, r *http.Request) {
	// Implement component data validation logic
}
