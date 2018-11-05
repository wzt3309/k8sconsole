package security

import "github.com/wzt3309/k8sconsole/src/app/backend/api"

type (
	// RequestBouncer represents an entity that manages API request access
	RequestBouncer struct {
		jwtService            api.JWTService
		userService           api.UserService
		teamMembershipService api.TeamMembershipService
		authDisabled          bool
	}

	// RequestBouncerParams represents the required parameters to create a new RequestBouncer instance
	RequestBouncerParams struct {
		JWTService            api.JWTService
		UserService           api.UserService
		TeamMembershipService api.TeamMembershipService
		AuthDisabled          bool
	}

	// RestrictedRequestContext is a data structure containing information used in RestrictedAccess
	RestrictedRequestContext struct {
		IsAdmin         bool
		IsTeamLeader    bool
		UserID          api.UserID
		UserMemberships []api.TeamMembership
	}
)
