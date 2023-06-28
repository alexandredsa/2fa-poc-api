package main

import (
	"log"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/repositories"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers"
	"github.com/alexandredsa/2fa-poc-api/pkg/config"
	"github.com/alexandredsa/2fa-poc-api/pkg/http"
)

func main() {
	// Load the database configuration
	dbConfig, err := config.LoadDatabaseConfig()
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
	}

	// Establish the database connection
	db, err := config.NewDatabaseConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	userRepository := repositories.NewUserRepository(db)
	tokenRepository := repositories.NewTokenRepository()

	authService := services.NewAuthenticationService(*userRepository, *tokenRepository)
	componentService := services.NewComponentService(*userRepository)
	userService := services.NewUserService(*userRepository)

	authHandler := handlers.NewAuthHandler(authService, componentService)
	accountDataHandler := handlers.NewAccountDataHandler(userService)
	twofaDataHandler := handlers.NewTwoFADataHandler(componentService)

	router := http.NewRouter(authHandler, accountDataHandler, twofaDataHandler)

	server := http.NewServer(":8080", router)

	log.Println("Server started on port 8080")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
