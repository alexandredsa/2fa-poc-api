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
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByID(userID string) (*models.User, error) {
	user := &models.User{}
	err := r.db.Where("id = ?", userID).First(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
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
