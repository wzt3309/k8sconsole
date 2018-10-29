package auth

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	authApi "github.com/wzt3309/k8sconsole/src/app/backend/auth/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
)

type frontendAuthManager struct {
	cryptoService           api.CryptoService
	jwtService              api.JWTService
	userService             api.UserService
	authenticationSkippable bool
}

func (self *frontendAuthManager) Login(payload *authApi.FrontendAuthPayload) (*authApi.FrontendAuthResponse, error) {
	if self.authenticationSkippable {
		return nil, errors.ErrAuthDisabled
	}

	u, err := self.userService.UserByUsername(payload.Username)
	if err != nil {
		return nil, err
	}

	err = self.authenticateInternal(u, payload.Password)
	if err != nil {
		return nil, err
	}

	token, err := self.generateToken(u)
	if err != nil {
		return nil, err
	}

	return &authApi.FrontendAuthResponse{
		JWTToken: token,
	}, nil
}

func (self *frontendAuthManager) authenticateInternal(user *api.User, password string) error {
	err := self.cryptoService.Verify(user.Password, password)
	if err != nil {
		return errors.ErrUnauthorized
	}
	return nil
}

func (self *frontendAuthManager) generateToken(user *api.User) (string, error) {
	tokenData := &api.TokenData{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
	}

	token, err := self.jwtService.Generate(tokenData)
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewFrontendAuthManager(cryptoService api.CryptoService, jwtService api.JWTService,
	userService api.UserService, authenticationSkippable bool) authApi.FrontendAuthManager {
	return &frontendAuthManager{
		cryptoService:           cryptoService,
		jwtService:              jwtService,
		userService:             userService,
		authenticationSkippable: authenticationSkippable,
	}
}
