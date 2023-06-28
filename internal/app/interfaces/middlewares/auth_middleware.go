package middlewares

import (
	"net/http"
)

// RequireAuthentication is a middleware that checks if the request is authenticated.
// If the request is not authenticated, it returns a 401 Unauthorized response.
func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Implement your authentication logic here
		// For example, check if the user is authenticated based on the request's authentication token
		// If the user is not authenticated, return a 401 Unauthorized response

		// If the user is authenticated, call the next handler
		next.ServeHTTP(w, r)
	})
}
