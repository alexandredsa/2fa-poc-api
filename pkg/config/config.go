package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DatabaseConfig holds the configuration parameters for the database connection.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// AppRepository
type AppRepository interface {
	Migrate(db *gorm.DB) error
}

// LoadDatabaseConfig loads the database configuration parameters from environment variables.
func LoadDatabaseConfig() (*DatabaseConfig, error) {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Validate required parameters
	if host == "" || port == "" || user == "" || dbName == "" {
		return nil, fmt.Errorf("missing required database configuration parameters")
	}

	return &DatabaseConfig{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DBName:   dbName,
	}, nil
}

// NewDatabaseConnection establishes a new database connection using the provided configuration.
func NewDatabaseConnection(config *DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}

	return db, nil
}

// MigrateAll performs database migrations for the given repositories
// using the provided gorm.DB connection.
func MigrateAll(db *gorm.DB, repositories []AppRepository) error {
	for _, r := range repositories {
		if err := r.Migrate(db); err != nil {
			return err
		}
	}

	return nil
}
