package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/repositories"
	"github.com/alexandredsa/2fa-poc-api/internal/app/domain/services"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers/account"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers/auth"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers/twofa"
	"github.com/alexandredsa/2fa-poc-api/pkg/applog"
	"github.com/alexandredsa/2fa-poc-api/pkg/appredis"
	"github.com/alexandredsa/2fa-poc-api/pkg/config"
	"github.com/alexandredsa/2fa-poc-api/pkg/http"
	"github.com/alexandredsa/2fa-poc-api/pkg/notification/notifier"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger := applog.NewLogger("main")
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
	if err := config.MigrateAll(db, []config.AppRepository{userRepository}); err != nil {
		log.Fatalf("Failed to setup database: %v", err)
	}

	redisClient, err := appredis.NewRedisClient()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	notifierFactory := notifier.NewNotifierFactory(redisClient)

	authService := services.NewAuthenticationService(*userRepository, notifierFactory)
	componentService := services.NewComponentService(*userRepository)
	userService := services.NewUserService(*userRepository)

	authHandler := auth.NewHandler(authService, componentService)
	accountDataHandler := account.NewHandler(userService)
	twofaDataHandler := twofa.NewHandler(componentService)

	router := http.NewRouter(authHandler, accountDataHandler, twofaDataHandler)

	server := http.NewServer(":8080", router)

	logger.Info("Server started on port 8080")
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
