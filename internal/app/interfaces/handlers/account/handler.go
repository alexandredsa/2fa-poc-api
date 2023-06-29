package account

import (
	"net/http"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
)

// Handler handles the account data-related HTTP requests.
type Handler struct {
	UserService *services.UserService
}

// NewHandler creates a new instance of the Handler.
func NewHandler(userService *services.UserService) *Handler {
	return &Handler{
		UserService: userService,
	}
}

// UpdateCredentials handles the update of user credentials.
func (h *Handler) UpdateCredentials(w http.ResponseWriter, r *http.Request) {
	// Implement update credentials logic
}
