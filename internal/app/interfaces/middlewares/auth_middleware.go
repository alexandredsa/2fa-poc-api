package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/alexandredsa/2fa-poc-api/pkg/httputils"
	"github.com/dgrijalva/jwt-go"
)

type ClaimKey string

const (
	ClaimUserID ClaimKey = "userID"
)

// RequireAuthentication is a middleware that checks if the request is authenticated.
// If the request is not authenticated, it returns a 401 Unauthorized response.
func RequireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the JWT token from the request header
		tokenString := httputils.ExtractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Missing authentication token", http.StatusUnauthorized)
			return
		}

		// Validate the JWT token
		claims, err := ValidateJWT(tokenString, []byte(os.Getenv("JWT_SECRET")))
		if err != nil {
			http.Error(w, "Invalid authentication token", http.StatusUnauthorized)
			return
		}

		// Retrieve the component from the JWT claims
		userID, ok := claims["user_id"].(string)
		if !ok {
			http.Error(w, "Invalid component in authentication token", http.StatusUnauthorized)
			return
		}

		// Set the component value in the request context
		ctx := context.WithValue(r.Context(), ClaimUserID, userID)
		r = r.WithContext(ctx)

		// If the user is authenticated, call the next handler
		next.ServeHTTP(w, r)
	})
}

// ValidateJWT validates the JWT token and returns the claims.
func ValidateJWT(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	// Check if the token is valid
	if !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	// Extract the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to extract JWT claims")
	}

	return claims, nil
}
