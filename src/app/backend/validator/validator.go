package validator

import (
	va "github.com/asaskevich/govalidator"
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
)

func IsUsername(str string) bool {
	if va.IsNull(str) || va.Contains(str, " ") {
		return false
	}
	return true
}

func IsRole(i interface{}, context interface{}) bool {
	switch v := i.(type) {
	case int:
		role := api.UserRole(v)
		if role != api.AdminRole && role != api.NormalUserRole {
			return false
		}
		return true
	default:
		return false
	}
}
