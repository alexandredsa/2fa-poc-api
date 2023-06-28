package middlewares

import (
	"net/http"
)

// TwoFAMiddleware is a middleware to enforce 2FA authentication.
func TwoFAMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement 2FA middleware logic

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
