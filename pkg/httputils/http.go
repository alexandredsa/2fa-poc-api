package httputils

import (
	"encoding/json"
	"net/http"
	"strings"
)

func WriteJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// ExtractTokenFromHeader extracts the JWT token from the Authorization header.
func ExtractTokenFromHeader(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	if authorizationHeader == "" {
		return ""
	}

	// Expect the Authorization header value to be in the format "Bearer <token>"
	tokenParts := strings.Split(authorizationHeader, " ")
	if len(tokenParts) != 2 || strings.ToLower(tokenParts[0]) != "bearer" {
		return ""
	}

	return tokenParts[1]
}
