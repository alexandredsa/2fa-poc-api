package repositories

import (
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"
	"gorm.io/gorm"
)

// UserRepository handles user data operations.
type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// CreateUser creates a new user in the database.
func (r *UserRepository) CreateUser(user *models.User) error {
	// Implement logic to save the user in the database using the provided db connection and GORM
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}
