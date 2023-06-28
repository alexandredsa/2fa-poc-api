package persistence

import (
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
)

// TokenRepository is an implementation of the TokenRepository interface.
type TokenRepository struct {
	// Add any necessary dependencies or database connections here
}

// SaveToken saves an authentication token in the persistent storage.
func (r *TokenRepository) SaveToken(token *models.Token) error {
	// Implement token saving in persistent storage
	return nil
}
