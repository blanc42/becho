package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idLength = 11
)

func GenerateShortID() (string, error) {
	bytes := make([]byte, idLength)
	alphabetLength := big.NewInt(int64(len(alphabet)))

	for i := 0; i < idLength; i++ {
		n, err := rand.Int(rand.Reader, alphabetLength)
		if err != nil {
			return "", err
		}
		bytes[i] = alphabet[n.Int64()]
	}

	return string(bytes), nil
}
