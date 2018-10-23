package file

// FileService represents a service for managing files
type FileService interface {
	FileExists(path string) (bool, error)
}
