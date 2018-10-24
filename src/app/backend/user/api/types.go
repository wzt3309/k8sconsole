package api

type (

	// User represents a user account
	User struct {
		ID       UserID   `json:"Id"`
		Username string   `json:"Username"`
		Password string   `json:"Password,omitempty"`
		Role     UserRole `json:"Role"`
	}

	// UserID represents a user identifier
	UserID int

	// UserRole represents the role of a user. It can be either an admin or a regular user
	UserRole int

	// UserService is used for user manager
	UserService interface {
		User(UserID) (*User, error)
		UserByUsername(string) (*User, error)
		Users() ([]User, error)
		UsersByRole(UserRole) ([]User, error)
		UpdateUser(UserID, *User) error
		CreateUser(*User) error
		DeleteUser(UserID) error
	}
)
