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

func (r *UserRepository) Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{})
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

// GetUserByUsername retrieves a user by their username from the database.
func (r *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, err
	}
	return &user, nil
}
