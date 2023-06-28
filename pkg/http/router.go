package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/handlers"
	"github.com/alexandredsa/2fa-poc-api/internal/app/interfaces/middlewares"
)

// NewRouter creates a new instance of the HTTP router.
func NewRouter(
	authHandler *handlers.AuthHandler,
	accountDataHandler *handlers.AccountDataHandler,
	twoFADataHandler *handlers.TwoFADataHandler,
) http.Handler {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))
	router.Use(middlewares.TwoFAMiddleware)

	// Auth routes
	router.Route("/auth", func(authRouter chi.Router) {
		authRouter.Post("/register", authHandler.Register)
		authRouter.Post("/login", authHandler.Login)
		authRouter.Post("/2fa/{component}/request", authHandler.RequestTwoFA)
		authRouter.Post("/2fa/{component}/validate", authHandler.ValidateTwoFA)
		authRouter.Put("/credentials", authHandler.UpdateCredentials)
		authRouter.Put("/{component}", authHandler.UpdateComponentData)
		authRouter.Post("/{component}/validate", authHandler.ValidateComponentData)
	})

	// Account data routes
	router.Route("/account", func(accountRouter chi.Router) {
		accountRouter.Put("/credentials", accountDataHandler.UpdateCredentials)
	})

	// TwoFA data routes
	router.Route("/2fa", func(twoFARouter chi.Router) {
		twoFARouter.Put("/{component}", twoFADataHandler.UpdateComponentData)
		twoFARouter.Post("/{component}/validate", twoFADataHandler.ValidateComponentData)
	})

	return router
}
