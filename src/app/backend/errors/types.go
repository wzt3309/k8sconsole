package errors

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

type Error string

func (e Error) Error() string {
	return string(e)
}
