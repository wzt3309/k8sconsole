package errors

import "github.com/emicklei/go-restful"

// General errors.
const (
	ErrUnauthorized   = Error("Unauthorized")
	ErrObjectNotFound = Error("Object not found inside the database")
	// server started with auth disabled
	ErrAuthDisabled = Error("Authentication is disabled")
)

// User errors.
const (
	ErrUserAlreadyExists = Error("User already exists")
)

// JWT errors.
const (
	ErrSecretGeneration = Error("Unable to generate secret key")
	ErrInvalidJWTToken  = Error("Invalid JWT token")
)

// Crypto errors.
const (
	ErrCryptoHashFailure = Error("Unable to hash data")
)

// Handle restful api route function error
type (
	// LoggerHandler defines a route function that includes a HandlerError return pointer
	LoggerHandler func(*restful.Request, *restful.Response) *HandlerError

	// HandlerError represents an error raised inside a route function
	HandlerError struct {
		StatusCode int
		Message    string
		Err        error
	}

	errorResponse struct {
		Err     string `json:"err,omitempty"`
		Details string `json:"details,omitempty"`
	}
)

type Error string

func (e Error) Error() string {
	return string(e)
}
