package mbird

type VerificationStatus string

const (
	StatusSent     VerificationStatus = "sent"
	StatusExpired  VerificationStatus = "expired"
	StatusFailed   VerificationStatus = "failed"
	StatusVerified VerificationStatus = "verified"
	StatusDeleted  VerificationStatus = "deleted"
)
