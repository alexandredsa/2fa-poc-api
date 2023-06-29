package auth

import "github.com/alexandredsa/2fa-poc-api/internal/app/domain/models"

// RegisterResponse represents the response body for user registration.
type RegisterResponse struct {
	User struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	} `json:"user"`
	Message string `json:"message"`
}

func NewRegisterResponse(user models.User, message string) RegisterResponse {
	r := RegisterResponse{}
	r.User.Name = user.Name
	r.User.Username = user.Username
	r.User.Password = user.Password
	r.User.Email = user.Email
	r.User.Phone = user.Phone
	r.Message = message

	return r
}
