package user

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"github.com/wzt3309/k8sconsole/src/app/backend/errors"
	userApi "github.com/wzt3309/k8sconsole/src/app/backend/user/api"
)

type userManager struct {
	userService   api.UserService
	cryptoService api.CryptoService
}

// UserCreate create user
func (self *userManager) UserCreate(payload *userApi.UserCreatePayload) (*api.User, error) {
	err := payload.Validate()
	if err != nil {
		return nil, err
	}

	user, err := self.userService.UserByUsername(payload.Username)
	if err != nil && err != errors.ErrObjectNotFound {
		return nil, err
	}

	if user != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	user = &api.User{
		Username: payload.Username,
		Role:     api.UserRole(payload.Role),
	}

	user.Password, err = self.cryptoService.Hash(payload.Password)
	if err != nil {
		return nil, errors.ErrCryptoHashFailure
	}

	err = self.userService.CreateUser(user)
	if err != nil {
		return nil, err
	}

	userApi.HideFields(user)

	return user, nil
}

func (self *userManager) UserList() ([]api.User, error) {
	users, err := self.userService.Users()
	if err != nil {
		return nil, err
	}

	for idx := range users {
		userApi.HideFields(&users[idx])
	}

	return users, nil
}

func NewUserManager(userService api.UserService, cryptoService api.CryptoService) userApi.UserManager {
	return &userManager{
		userService:   userService,
		cryptoService: cryptoService,
	}
}
