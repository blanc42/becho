package utils

import (
	"crypto/rand"
	"encoding/base64"
)

func GenerateShortID() (string, error) {
	// Generate 16 random bytes (128 bits, same as UUID)
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Encode to base64
	encoded := base64.RawURLEncoding.EncodeToString(b)

	// Return the first 11 characters
	return encoded[:11], nil
}
