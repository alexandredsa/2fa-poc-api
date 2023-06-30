package auth_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/repositories"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers/auth"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification/notifier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDatabase() (*gorm.DB, error) {
	// Connect to the SQLite in-memory database
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto-migrate the User model to create the "users" table
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	// Create a user record
	user := models.User{
		ID:       "userID",
		Username: "testuser",
		Password: "testpassword",
	}

	// Insert the user record into the "users" table
	err = db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestRegisterUserHandler(t *testing.T) {
	db, err := setupTestDatabase()
	require.NoError(t, err)
	// Create a sample user payload
	userPayload := map[string]string{
		"username": "john.doe",
		"password": "password123",
		"email":    "john.doe@example.com",
	}

	// Convert user payload to JSON
	payloadBytes, _ := json.Marshal(userPayload)
	body := bytes.NewReader(payloadBytes)

	// Create a new HTTP request
	req := httptest.NewRequest("POST", "/register", body)
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder to capture the handler response
	recorder := httptest.NewRecorder()

	userRepository := *repositories.NewUserRepository(db)
	// Create an instance of the mock authentication service
	authService := services.NewAuthenticationService(userRepository, notifier.NotifierFactory{})

	// Create the server handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Instantiate the handler with the mock authentication service
		authHandler := auth.NewHandler(authService, nil)

		// Call the RegisterUser handler function
		authHandler.Register(w, r)
	}

	// Serve the HTTP request
	handler(recorder, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Parse the response body
	var response auth.RegisterResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Access the response data
	message := response.Message

	// Assert the expected response
	expectedMessage := "Registration successful"
	assert.Equal(t, expectedMessage, message)
}
