package rest

import (
	"net/http"
	"os"

	commonAuth "github.com/emitra-labs/common/auth"
	"github.com/emitra-labs/common/rest/handler"
	"github.com/emitra-labs/common/rest/server"
	"github.com/emitra-labs/identity-service/controller"
	"github.com/emitra-labs/identity-service/controller/session"
	"github.com/emitra-labs/identity-service/controller/user"
	"github.com/emitra-labs/identity-service/controller/verification"
	"github.com/swaggest/openapi-go/openapi31"
)

func NewServer() *server.Server {
	description := "User management and authentication service."
	spec := openapi31.Spec{
		Openapi: "3.1.0",
		Info: openapi31.Info{
			Title:       "Identity Service",
			Version:     "0.0.1",
			Description: &description,
		},
		Servers: []openapi31.Server{
			{URL: os.Getenv("OPENAPI_SERVER_URL")},
		},
	}

	// Parse JWT public key
	jwtPublicKey, err := commonAuth.ParsePublicKeyFromBase64(os.Getenv("BASE64_JWT_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	// Create new server
	s := server.New(server.Config{
		OpenAPI:      server.OpenAPI{Spec: &spec},
		JWTPublicKey: jwtPublicKey,
	})

	handler.AddHealthCheck(s)

	// Auth
	handler.Add(s, http.MethodPost, "/auth/sign-up", controller.SignUp, handler.Config{
		Summary:     "Sign up",
		Description: "Signing up for a new user account. The user need to verify their email afterward.",
		Tags:        []string{"Auth"},
	})
	handler.Add(s, http.MethodPost, "/auth/sign-in", controller.SignIn, handler.Config{
		Summary:     "Sign in",
		Description: "Sign in",
		Tags:        []string{"Auth"},
	})
	handler.Add(s, http.MethodPost, "/auth/refresh", controller.RefreshToken, handler.Config{
		Summary:     "Refresh token",
		Description: "Refreshing the token will generate new tokens. Typically, you will need to do this when your access token has expired.",
		Tags:        []string{"Auth"},
	})
	handler.Add(s, http.MethodPost, "/auth/sign-out", controller.SignOut, handler.Config{
		Summary:      "Sign out",
		Description:  "Signing out will delete the current user session, so refreshing the token will no longer be possible.",
		Tags:         []string{"Auth"},
		Authenticate: true,
	})

	// Session
	handler.Add(s, http.MethodGet, "/users/current/sessions", session.GetCurrentUserSessions, handler.Config{
		Summary:      "Get current user's sessions",
		Description:  "Get current user's sessions",
		Tags:         []string{"Session"},
		Authenticate: true,
	})
	handler.Add(s, http.MethodGet, "/users/current/sessions/current", session.GetCurrentUserCurrentSession, handler.Config{
		Summary:      "Get current user's active session",
		Description:  "Get current user's active session",
		Tags:         []string{"Session"},
		Authenticate: true,
	})
	handler.Add(s, http.MethodGet, "/users/current/sessions/:sessionId", session.GetCurrentUserSession, handler.Config{
		Summary:      "Get current user's session",
		Description:  "Get current user's session",
		Tags:         []string{"Session"},
		Authenticate: true,
	})

	// User
	handler.Add(s, http.MethodPost, "/users", user.CreateUser, handler.Config{
		Summary:     "Create user",
		Description: "Create a new user. You must be a super admin to access this resource.",
		Tags:        []string{"User"},
		SuperAdmin:  true,
	})
	handler.Add(s, http.MethodGet, "/users", user.GetUsers, handler.Config{
		Summary:     "Get users",
		Description: "Retrive all users. You must be a super admin to access this resource.",
		Tags:        []string{"User"},
		SuperAdmin:  true,
	})
	handler.Add(s, http.MethodGet, "/users/current", user.GetCurrentUser, handler.Config{
		Summary:      "Get current user",
		Description:  "Get current user",
		Tags:         []string{"User"},
		Authenticate: true,
	})
	handler.Add(s, http.MethodGet, "/users/:userId", user.GetUser, handler.Config{
		Summary:     "Get user",
		Description: "Retrive a user. You must be a super admin to access this resource.",
		Tags:        []string{"User"},
		SuperAdmin:  true,
	})
	handler.Add(s, http.MethodPatch, "/users/current", user.UpdateCurrentUser, handler.Config{
		Summary:      "Update current user",
		Description:  "Update current user",
		Tags:         []string{"User"},
		Authenticate: true,
	})
	handler.Add(s, http.MethodPatch, "/users/:userId", user.UpdateUser, handler.Config{
		Summary:     "Update a user",
		Description: "Update a user. You must be a super admin to access this resource.",
		Tags:        []string{"User"},
		SuperAdmin:  true,
	})
	handler.Add(s, http.MethodDelete, "/users/current", user.DeleteCurrentUser, handler.Config{
		Summary:      "Delete current user",
		Description:  "Delete current user",
		Tags:         []string{"User"},
		Authenticate: true,
	})
	handler.Add(s, http.MethodDelete, "/users/:userId", user.DeleteUser, handler.Config{
		Summary:     "Delete a user",
		Description: "Delete a user. You must be a super admin to access this resource.",
		Tags:        []string{"User"},
		SuperAdmin:  true,
	})

	// Verification
	handler.Add(s, http.MethodPost, "/verify", verification.Verify, handler.Config{
		Summary:     "Verify user",
		Description: "Verify user using data taken from verification email",
		Tags:        []string{"Verification"},
	})

	return s
}
