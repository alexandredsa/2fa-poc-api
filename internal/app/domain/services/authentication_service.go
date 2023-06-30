package services

import (
	"context"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/repositories"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification/notifier"
	"github.com/google/uuid"     // for generating random IDs
	"golang.org/x/crypto/bcrypt" // for hashing passwords
)

// AuthenticationService represents a service for authentication-related operations.
type AuthenticationService struct {
	userRepository  repositories.UserRepository
	notifierFactory notifier.NotifierFactory
}

func NewAuthenticationService(userRepository repositories.UserRepository, notifierFactory notifier.NotifierFactory) *AuthenticationService {
	return &AuthenticationService{
		userRepository:  userRepository,
		notifierFactory: notifierFactory,
	}
}

// RegisterUser registers a new user and returns the created user.
func (s *AuthenticationService) RegisterUser(user *models.User) (*models.User, error) {
	// Generate a random ID for the user
	user.ID = uuid.New().String()

	// Hash the user's password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Set the hashed password for the user
	user.Password = string(hashedPassword)

	// Save the user in the database
	err = s.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUserByUsername retrieves a user by their username.
func (s *AuthenticationService) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		// Handle any errors that occurred during the retrieval
		return nil, err
	}

	return user, nil
}

// AuthenticateUser performs user authentication.
func (s *AuthenticationService) AuthenticateUser(loginRequest *models.LoginRequest) (*models.Token, error) {
	// Implement user authentication logic
	return nil, nil
}

// RequestTwoFACode requests a 2FA code for a given component.
func (s *AuthenticationService) RequestTwoFACode(ctx context.Context, userID string, component string) error {
	// Implement 2FA code request logic
	notifier := s.notifierFactory.NewNotifier(component)

	user, err := s.userRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	return notifier.SendVerificationCode(ctx, *user)
}

func (s *AuthenticationService) SendCodeConfirmation(uuserID string, component string) (*models.Token, error) {
	// Implement 2FA code validation logic
	return nil, nil
}
