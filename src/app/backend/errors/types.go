package errors

const (
	ErrObjectNotFound = Error("Object not found inside the database")
	// server started with auth disabled
	ErrAuthDisabled = Error("Authentication is disabled")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
