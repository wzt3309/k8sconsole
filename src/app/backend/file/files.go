package file

import "os"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// FileExists checks for the existence of the specified file.
func (service *Service) FileExists(filePath string) (bool, error) {
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
