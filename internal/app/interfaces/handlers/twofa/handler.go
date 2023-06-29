package twofa

import (
	"net/http"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
	"github.com/go-chi/chi"
)

// Handler handles the 2FA data-related HTTP requests.
type Handler struct {
	ComponentService *services.ComponentService
}

// NewHandler creates a new instance of the Handler.
func NewHandler(componentService *services.ComponentService) *Handler {
	return &Handler{
		ComponentService: componentService,
	}
}

// UpdateComponentData handles the update of component data.
func (h *Handler) UpdateComponentData(w http.ResponseWriter, r *http.Request) {
	// Extract the component from the URL path params
	component := chi.URLParam(r, "component")

	// Implement update component data logic based on the component
	switch component {
	case "component1":
		// Update component1 data
	case "component2":
		// Update component2 data
	default:
		// Invalid component
	}
}

// ValidateComponentData handles the validation of component data.
func (h *Handler) ValidateComponentData(w http.ResponseWriter, r *http.Request) {
	// Extract the component from the URL path params
	component := chi.URLParam(r, "component")

	// Implement component data validation logic based on the component
	switch component {
	case "component1":
		// Validate component1 data
	case "component2":
		// Validate component2 data
	default:
		// Invalid component
	}
}
