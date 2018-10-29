package jwt

import (
	"github.com/stretchr/testify/assert"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"testing"
)

func TestJWTService(t *testing.T) {

	jwtSVC, err := NewJWTService()
	if err != nil {
		t.Fatal(err)
	}

	expected := &api.TokenData{
		ID:       api.UserID(1),
		Username: "john",
		Role:     api.UserRole(1),
	}

	signedToken, err := jwtSVC.Generate(expected)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(signedToken)

	actual, err := jwtSVC.Decrypt(signedToken)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(actual)
	assert.Equal(t, expected, actual)
}
