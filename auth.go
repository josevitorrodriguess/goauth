// auth.go
package auth

import (
	"context"
	"net/http"
	"time"
)

// AuthenticatedUser represents basic informations for authenticated user
type AuthenticatedUser struct {
	ID       string
	Username string
	Roles    []string
	// Additional fields that may be useful in different scenarios
	Metadata map[string]interface{} // Map of additional data that applications may need to store (such as preferences, access level, etc.)
}

type Authenticator interface {
	// Authenticate processes a request and extracts/validates credentials
	Authenticate(ctx context.Context, r *http.Request) (*AuthenticatedUser, error)

	// GenerateCredentials creates credentials for the user (token, cookie, etc.)
	GenerateCredentials(user *AuthenticatedUser) (interface{}, error)

	// ValidateCredentials checks if credentials are valid
	ValidateCredentials(creds interface{}) (*AuthenticatedUser, error)

	// WriteCredentials write creddentials in the HTTP response (headers, cookies, etc.)
	WriteCredentials(w http.ResponseWriter, creds interface{}) error

	// Middleware return middleware  HTTP for authentication
	Middleware() func(http.Handler) http.Handler

	// MiddlewareWithRoles return middleware  retorna um middleware which also checks roles/permissions
	MiddlewareWithRoles(roles ...string) func(http.Handler) http.Handler
}

// AuthOptions general settings for any authenticator
type AuthOptions struct {
	// standard duration for credentials
	CredentialsDuration time.Duration

	// Handler for when authenticacion fails
	UnauthorizedHandler http.Handler

	// Handler for when access is denied (role verification failed)
	ForbiddenHandler http.Handler

	// Func for extract crendentials for  request
	CredentialsExtractor func(r *http.Request) (interface{}, error)
}

type AuthError struct {
	Code    int
	Message string
	Err     error
}

func (e AuthError) Error() string {
	return e.Message
}

func (e AuthError) Unwrap() error {
	return e.Err
}

var (
	ErrInvalidCredentials = AuthError{Code: http.StatusUnauthorized, Message: "invalid credentials"}
	ErrExpiredCredentials = AuthError{Code: http.StatusUnauthorized, Message: "expired credentials"}
	ErrMissingCredentials = AuthError{Code: http.StatusUnauthorized, Message: "crendentials not provided"}
	ErrAccessDenied       = AuthError{Code: http.StatusForbidden, Message: "access denied"}
)

// UserFromContext extract a authenticated user from context
func UserFromContext(ctx context.Context) (*AuthenticatedUser, bool) {
	user, ok := ctx.Value(userContextKey).(*AuthenticatedUser)
	return user, ok
}

// SetUserContext add a athenticated user in context
func SetUserContext(ctx context.Context, user *AuthenticatedUser) context.Context {
	return context.WithValue(ctx, userContextKey, user)
}

// private key for context
type contextKey string

const userContextKey contextKey = "authenticated_user"
