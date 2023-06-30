package notification

import (
	"crypto/rand"
	"math/big"
)

type Generator interface {
	// NewCode uses a random algorithm to create an alphanumeric value ([A-Z][0-9]) of the defined length
	NewCode(length int) string
}

type BasicGenerator struct{}

func NewGenerator() Generator {
	return &BasicGenerator{}
}

func (g *BasicGenerator) NewCode(length int) string {
	const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	code := make([]byte, length)
	max := big.NewInt(int64(len(charset)))

	for i := 0; i < length; i++ {
		index, _ := rand.Int(rand.Reader, max)
		code[i] = charset[index.Int64()]
	}

	return string(code)
}
