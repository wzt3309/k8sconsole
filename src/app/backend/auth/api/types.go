package api

// FrontendAuthPayload contains 'username' and 'password' of user.
// It uses struct field tags to validate fields by "govalidator".
type FrontendAuthPayload struct {
	Username string `valid:"runelength(2|20),required"`
	Password string `valid:"runelength(6|20),required"`
}

// FrontendAuthResponse returned from our backend as a response for login request.
// It contains generated JWTToken and a list of errors during th authentication.
type FrontendAuthResponse struct {
	JWTToken string `json:"jwtToken"`
}

// FrontendAuthManager is used for user authentication manager.
type FrontendAuthManager interface {
	Login(*FrontendAuthPayload) (*FrontendAuthResponse, error)
}
