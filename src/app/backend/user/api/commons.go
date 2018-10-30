package api

import (
	"github.com/asaskevich/govalidator"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
)

func (payload *UserCreatePayload) Validate() error {
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return err
	}
	return nil
}

func HideFields(user *api.User) {
	user.Password = ""
}
