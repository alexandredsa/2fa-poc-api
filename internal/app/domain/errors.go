package domain

import "errors"

// Custom errors for your domain

var (
	ErrNotFound       = errors.New("not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrInvalidPayload = errors.New("invalid payload")
	// Add any other custom errors specific to your domain
)
