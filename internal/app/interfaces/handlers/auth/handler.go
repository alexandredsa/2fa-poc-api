package auth

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/middlewares"
	"github.com/alexandredsa/2fa-poc-api/pkg/httputils"
	"github.com/go-chi/chi"

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

	// Hash the user's password before storing it
	err = HashPassword(user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Call the AuthService to create the user
	createdUser, err := h.authService.RegisterUser(user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// Prepare the response
	response := NewRegisterResponse(*createdUser, "Registration successful")

	// Write the response
	httputils.WriteJSONResponse(w, response)
}

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
	err = ComparePasswords(user.Password, requestBody.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	tokenString, err := GenerateJWTToken(user.ID)
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

	// Write the response
	httputils.WriteJSONResponse(w, response)
}

func HashPassword(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func ComparePasswords(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func GenerateJWTToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"iss": "2fa-poc-api",
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	})

	secretKey := []byte(os.Getenv("JWT_SECRET"))

	return token.SignedString(secretKey)
}

// RequestTwoFA handles the request for 2FA code.
func (h *Handler) RequestTwoFA(w http.ResponseWriter, r *http.Request) {
	// Extract the component from the URL parameter
	component := chi.URLParam(r, "component")

	ctx := r.Context()
	// Get the user ID from the request context
	// (assuming it has been set during authentication)
	userID := ctx.Value(middlewares.ClaimUserID).(string)

	if err := h.authService.RequestTwoFACode(ctx, userID, component); err != nil {
		http.Error(w, "Failed to send email 2FA request", http.StatusInternalServerError)
		return
	}

	// Return a success response
	response := map[string]string{
		"message": "2FA request successful",
	}

	httputils.WriteJSONResponse(w, response)
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
