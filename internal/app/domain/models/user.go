package models

// User represents a user in the system.
type User struct {
	ID       int
	Username string
	Password string
	Email    string
	Phone    string
}
