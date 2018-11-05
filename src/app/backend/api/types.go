package api

const (
	_ UserRole = iota
	AdminRole
	NormalUserRole
)

const (
	_ ResourceAccessLevel = iota
	// ReadWriteAccessLevel represents an access level with read-write permissions on a resource
	ReadWriteAccessLevel
)

// APIs about user management
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

// APIs about team
type (
	// Team represents a list of user accounts
	Team struct {
		ID   TeamID `json:"Id"`
		Name string `json:"Name"`
	}

	// TeamID represents a team identifier
	TeamID int

	// TeamMembership represents a membership association between a user and a team
	TeamMembership struct {
		ID     TeamMembershipID `json:"Id"`
		UserID UserID           `json:"UserID"`
		TeamID TeamID           `json:"TeamID"`
		Role   MembershipRole   `json:"Role"`
	}

	// TeamMembershipID represents a team membership identifier
	TeamMembershipID int

	// MembershipRole represents the role of a user within a team
	MembershipRole int

	// TeamService represents a service for managing user data
	TeamService interface {
		Team(ID TeamID) (*Team, error)
		TeamByName(name string) (*Team, error)
		Teams() ([]Team, error)
		CreateTeam(team *Team) error
		UpdateTeam(ID TeamID, team *Team) error
		DeleteTeam(ID TeamID) error
	}

	// TeamMembershipService represents a service for managing team membership data
	TeamMembershipService interface {
		TeamMembership(ID TeamMembershipID) (*TeamMembership, error)
		TeamMemberships() ([]TeamMembership, error)
		TeamMembershipsByUserID(userID UserID) ([]TeamMembership, error)
		TeamMembershipsByTeamID(teamID TeamID) ([]TeamMembership, error)
		CreateTeamMembership(membership *TeamMembership) error
		UpdateTeamMembership(ID TeamMembershipID, membership *TeamMembership) error
		DeleteTeamMembership(ID TeamMembershipID) error
		DeleteTeamMembershipByUserID(userID UserID) error
		DeleteTeamMembershipByTeamID(teamID TeamID) error
	}
)

// APIs about authentication
type (

	// TokenData represents the data embedded in a JWT token
	TokenData struct {
		ID       UserID
		Username string
		Role     UserRole
	}

	// CryptoService represents a fileService for encrypting/hashing data
	CryptoService interface {
		Hash(data string) (string, error)
		Verify(hash string, data string) error
	}

	// JWTService represents a fileService for managing JWT tokens
	JWTService interface {
		// Generate token based on TokenData
		Generate(*TokenData) (string, error)
		// Verify and decrypt generated token
		Decrypt(string) (*TokenData, error)
	}
)

// APIs about files
type (

	// FileService represents a fileService for managing files
	FileService interface {
		FileExists(path string) (bool, error)
	}

	// DataStore represents a fileService for store information
	DataStore interface {
		Open() error
		Init() error
		Close() error
		GetUserService() UserService
	}
)

// Resource control
type (

	// ResourceControl represent a reference to a k8s resource with specific access controls
	ResourceControl struct {
		ID             ResourceControlID    `json:"Id"`
		ResourceID     string               `json:"ResourceId"`
		SubResourceIDs []string             `json:"SubResourceIds"`
		Type           ResourceControlType  `json:"Type"`
		UserAccesses   []UserResourceAccess `json:"UserAccesses"`
		TeamAccesses   []TeamResourceAccess `json:"TeamAccesses"`
		Public         bool                 `json:"Public"`
	}

	// ResourceControlID represents a resource control identifier
	ResourceControlID int

	// ResourceControlType represents the type of resource associated to the resource control (volume, pod, service...)
	ResourceControlType int

	// ResourceAccessLevel represents the level of control associated to a resource
	ResourceAccessLevel int

	// UserResourceAccess represents the level of control on a resource for a specific user
	UserResourceAccess struct {
		UserID      UserID              `json:"UserId"`
		AccessLevel ResourceAccessLevel `json:"AccessLevel"`
	}

	// TeamResourceAccess represents the level of control on a resource for a specific team
	TeamResourceAccess struct {
		TeamID      TeamID              `json:"TeamId"`
		AccessLevel ResourceAccessLevel `json:"AccessLevel"`
	}
)
