package errors

const (
	ErrObjectNotFound = Error("Object not found inside the database")
)

type Error string

func (e Error) Error() string {
	return string(e)
}