package api

import (
	"github.com/asaskevich/govalidator"
	"net/http"
)

// Validate the auth payload from frontend
func (payload *FrontendAuthPayload) Validate(r *http.Request) error {
	if _, err := govalidator.ValidateStruct(payload); err != nil {
		return err
	}
	return nil
}
