package api

import (
	"k8s.io/client-go/tools/clientcmd/api"
	"time"
)

const (
	// The name of key for encryption storage.
	EncrytionKeyHolderName       = "k8sconsole-key-holder"
	EncryptionKeyHolderNamespace = "kube-system"

	// Expiration time (in sec) of tokens generated by k8sconsole. Default: 15 min.
	DefaultTokenTTL = 900
)

// AuthenticationModes represents which auth mode supported by k8sconsole.
type AuthenticationModes map[AuthenticationMode]bool

// IsEnable returns true if the given auth mode is supported, false otherwise.
func (self AuthenticationModes) IsEnabled(mode AuthenticationMode) bool {
	_, exists := self[mode]
	return exists
}

// Array returns slice of auth modes supported by k8sconsole.
// The return value will be empty ([] not nil) if no mode supported.
func (self AuthenticationModes) Array() []AuthenticationMode {
	var modes []AuthenticationMode
	for mode := range self {
		modes = append(modes, mode)
	}

	if modes == nil {
		modes = []AuthenticationMode{}
	}

	return modes
}

// Add adds auth mode to AuthenticationModes map.
func (self AuthenticationModes) Add(mode AuthenticationMode) {
	self[mode] = true
}

// AuthenticationMode represents auth mode, i.e. basic(username and password).
type AuthenticationMode string

// String returns string representation of auth mode.
func (self AuthenticationMode) String() string {
	return string(self)
}

// Authentication modes supported by k8sconsole should be defined below.
const (
	Basic AuthenticationMode = "basic"
	Token AuthenticationMode = "token"
)

// AuthManager is used for user authentication management.
type AuthManager interface {
	// Login authenticates user based on provided LoginSpec and returns AuthResponse.
	Login(*LoginSpec) (*AuthResponse, error)
	// Refresh takes valid token that hasn't expired yet and returns a new one with expiration time set to TokenTTL.
	// In case provided token has expired, token expiration error is returned.
	Refresh(string) (string, error)
	// AuthenticationModes returns array of auth modes supported by k8sconsole.
	AuthenticationModes() []AuthenticationMode
	//  AuthenticationSkippable tells if the Skip button should be enabled or not
	AuthenticationSkippable() bool
}

// TokenManager is used for generated and decrypting tokens used for authorization.
// In this branch(quickv), Authorization is handled by k8s apiserver.
// Token contains AuthInfo used to create k8s api client
type TokenManager interface {
	// Generate secure token based on AuthInfo.
	Generate(api.AuthInfo) (string, error)
	// Decrypt generated token and extract AuthInfo from it which is used for creating k8s apiserver client
	Decrypt(string) (*api.AuthInfo, error)
	// Refresh returns refreshed token based on provided token. In case provided token has expired, token expiration
	// error is returned.
	Refresh(string) (string, error)
	// SetTokenTTL sets expiration time (in sec) of generated tokens.
	SetTokenTTL(time.Duration)
}

// Authenticator represents authentication methods, Currently supported types are:
//	- Basic 			- Username and password based authentication
//	- Token 			- Any bearer token accepted by apiserver
//	- kubeConfig 	- Authenticates user based on kubeconfig file.
type Authenticator interface {
	GetAuthInfo() (api.AuthInfo, error)
}

// LoginSpec is extracted from request coming from k8sconsole frontend during loging request. It contains all
// information required to authenticate user.
type LoginSpec struct {
	// Username is the username for 'basic' mode authentication.
	Username string `json:"username"`
	// Password is the password for 'basic' mode authentication.
	Password string `json:"password"`
	// Token is the jwe token for 'token' mode authentication.
	Token string `json:"token"`
	// KubeConfig is the content of users' kubeconfig file. We can extract all auth information
	// from the data in the file.
	KubeConfig string `json:"kubeConfig"`
}

// TokenRefreshSpec contains token that is required by token refresh operation.
type TokenRefreshSpec struct {
	// JWEToken is a token generated during login request that contains AuthInfo data in the payload.
	JWEToken string `json:"jweToken"`
}

// AuthResponse represents the response returned from k8sconsole backend for login requests. It contains generated
// JWEToken and a list of non-critical errors such as 'Failed authentication' to tell the frontend what unexpected
// happened during login request.
type AuthResponse struct {
	// JWEToken is a token generated during login request that contains AuthInfo data in the payload.
	JWEToken string `json:"jweToken"`
	// Errors are a list of non-critical errors that happened during login request.
	Errors []error `json:"errors"`
}

// LoginModesResponse contains list of auth modes supported by k8sconsole
type LoginModesResponse struct {
	Modes []AuthenticationMode `json:"modes"`
}

// LoginSkippableResponse contains a flag that tells the frontend not to display the 'auth skip' button
// It's just for hide the button, not disable unauthenticated access
type LoginSkippableResponse struct {
	Skippable bool `json:"skippable"`
}
