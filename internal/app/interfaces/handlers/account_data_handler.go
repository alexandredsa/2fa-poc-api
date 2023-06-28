package handlers

import (
	"net/http"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
)

// AccountDataHandler handles the account data-related HTTP requests.
type AccountDataHandler struct {
	UserService *services.UserService
}

// NewAccountDataHandler creates a new instance of the AccountDataHandler.
func NewAccountDataHandler(userService *services.UserService) *AccountDataHandler {
	return &AccountDataHandler{
		UserService: userService,
	}
}

// UpdateCredentials handles the update of user credentials.
func (h *AccountDataHandler) UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	// Implement update credentials logic
}
