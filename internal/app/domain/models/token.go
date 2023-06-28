package models

// Token represents an authentication token.
type Token struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int
	TwoFAEnabled bool
}
