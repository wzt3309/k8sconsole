package api

const (
	// BucketName represents the name of the bucket where Service store data
	BucketName = "users"
)

type (

	// User represents a user account
	User struct {
			ID UserID `json:"Id"`
			Username string `json:"Username"`
			Password string `json:"Password,omitempty"`
			Role UserRole `json:"Role"`
	}

	// UserID represents a user identifier
	UserID int

	// UserRole represents the role of a user. It can be either an admin or a regular user
	UserRole int
)