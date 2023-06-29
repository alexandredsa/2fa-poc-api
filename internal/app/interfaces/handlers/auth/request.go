package auth

// RegisterRequest represents the request body for user registration.
type RegisterRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// LoginRequest represents user credentials.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
