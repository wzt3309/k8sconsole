package api

import "github.com/wzt3309/k8sconsole/src/app/backend/api"

// UserCreatePayload represents payload to create a user
type UserCreatePayload struct {
	Username string `valid:"username,length(2|20)"`
	Password string `valid:"required,length(4|20)"`
	Role     int    `valid:"required,role"`
}

// UserManager is used to manage users
type UserManager interface {
	UserCreate(*UserCreatePayload) (*api.User, error)
	UserList() ([]api.User, error)
}
