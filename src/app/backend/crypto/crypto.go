package crypto

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"golang.org/x/crypto/bcrypt"
)

type service struct{}

func NewService() api.CryptoService {
	return &service{}
}

// Hash hashes a string using the bcrypt algorithm.
func (*service) Hash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

// Verify compares a hash to clear data and returns an error if the comparison fails.
func (*service) Verify(hash string, data string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
}
