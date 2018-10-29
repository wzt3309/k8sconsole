package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type Service struct{}

// Hash hashes a string using the bcrypt algorithm.
func (*Service) Hash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// Verify compares a hash to clear data and returns an error if the comparison fails.
func (*Service) Verify(hash string, data string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
}
