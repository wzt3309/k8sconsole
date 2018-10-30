package file

import (
	"github.com/wzt3309/k8sconsole/src/app/backend/api"
	"os"
)

type service struct{}

func NewService() api.FileService {
	return &service{}
}

// FileExists checks for the existence of the specified file.
func (*service) FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
