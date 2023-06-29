package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Handler handles the authentication-related HTTP requests.
type Handler struct {
	authService      *services.AuthenticationService
	componentService *services.ComponentService
}

// NewHandler creates a new instance of the Handler.
func NewHandler(authService *services.AuthenticationService, componentService *services.ComponentService) *Handler {
	return &Handler{
		authService:      authService,
		componentService: componentService,
	}
}

// Register handles the registration request.
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var request RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Create a new user based on the request data
	user := &models.User{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
		Phone:    request.Phone,
	}

	// Call the AuthService to create the user
	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}
	// Prepare the response
	response := NewRegisterResponse(*createdUser, "Registration successful")

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
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	var requestBody models.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Retrieve the user from the database by username
	user, err := h.authService.GetUserByUsername(requestBody.Username)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Compare the provided password with the stored hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"iss": "2fa-poc-api",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	// Sign the token with the secret key
	secretKey := []byte(os.Getenv("JWT_SECRET"))
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		http.Error(w, "Failed to generate access token", http.StatusInternalServerError)
		return
	}

	// Build the response
	response := LoginResponse{
		AccessToken:      tokenString,
		TokenType:        "Bearer",
		ExpiresIn:        time.Hour.Milliseconds() / 1000,
		TwoFAValidations: make([]string, 0), // Add the logic to retrieve 2FA validations
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Write the response JSON to the response writer
	json.NewEncoder(w).Encode(response)
}

// RequestTwoFA handles the request for 2FA code.
func (h *Handler) RequestTwoFA(w http.ResponseWriter, r *http.Request) {
	// Implement 2FA request logic
}

// ValidateTwoFA handles the validation of 2FA code.
func (h *Handler) ValidateTwoFA(w http.ResponseWriter, r *http.Request) {
	// Implement 2FA validation logic
}

// UpdateCredentials handles the update of user credentials.
func (h *Handler) UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	// Implement update credentials logic
}

// UpdateComponentData handles the update of component data.
func (h *Handler) UpdateComponentData(w http.ResponseWriter, r *http.Request) {
	// Implement update component data logic
}

// ValidateComponentData handles the validation of component data.
func (h *Handler) ValidateComponentData(w http.ResponseWriter, r *http.Request) {
	// Implement component data validation logic
}
