package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	"time"
)

// Service represents a service for JWTService
type Service struct {
	secret []byte
}

type claims struct {
	UserID   int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	jwt.StandardClaims
}

// NewService initializes a new service. It will generate a random key that will be used to sign JWT tokens.
func NewJWTService() (*Service, error) {
	secret := securecookie.GenerateRandomKey(32)
	if secret == nil {
		return nil, errors.ErrSecretGeneration
	}

	service := &Service{
		secret,
	}
	return service, nil
}

// Generate generates a new JWT Token
func (self *Service) Generate(data *api.TokenData) (string, error) {
	expireToken := time.Now().Add(time.Hour * 8).Unix()
	cl := claims{
		int(data.ID),
		data.Username,
		int(data.Role),
		jwt.StandardClaims{
			ExpiresAt: expireToken,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cl)

	signedToken, err := token.SignedString(self.secret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// Decrypt parses a JWT token and verify its validity. It returns an error if token is invalid.
func (self *Service) Decrypt(token string) (*api.TokenData, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			return nil, msg
		}
		return self.secret, nil
	})
	if err == nil && parsedToken != nil {
		if cl, ok := parsedToken.Claims.(*claims); ok && parsedToken.Valid {
			tokenData := &api.TokenData{
				ID:       api.UserID(cl.UserID),
				Username: cl.Username,
				Role:     api.UserRole(cl.Role),
			}
			return tokenData, nil
		}
	}
	return nil, errors.ErrInvalidJWTToken
}
