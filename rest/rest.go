package rest

import (
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	commonAuth "github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/server"
	"github.com/ukasyah-dev/identity-service/controller"
	"github.com/ukasyah-dev/identity-service/controller/auth"
	"github.com/ukasyah-dev/identity-service/controller/session"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/controller/verification"
)

var Server *server.Server

func init() {
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
	Server = server.New(server.Config{
		OpenAPI:      server.OpenAPI{Spec: &spec},
		JWTPublicKey: jwtPublicKey,
	})

	handler.Add(Server, http.MethodGet, "/", controller.HealthCheck, handler.Config{
		Summary:     "Health check",
		Description: "Check whether the server is ready to serve",
		Tags:        []string{"Health"},
	})

	// Auth
	handler.Add(Server, http.MethodPost, "/auth/sign-up", auth.SignUp, handler.Config{
		Summary:     "Sign up",
		Description: "Signing up for a new user account. The user need to verify their email afterward.",
		Tags:        []string{"Auth"},
	})
	handler.Add(Server, http.MethodPost, "/auth/sign-in", auth.SignIn, handler.Config{
		Summary:     "Sign in",
		Description: "Sign in",
		Tags:        []string{"Auth"},
	})
	handler.Add(Server, http.MethodPost, "/auth/refresh", auth.RefreshToken, handler.Config{
		Summary:     "Refresh token",
		Description: "Refreshing the token will generate new tokens. Typically, you will need to do this when your access token has expired.",
		Tags:        []string{"Auth"},
	})
	handler.Add(Server, http.MethodPost, "/auth/sign-out", auth.SignOut, handler.Config{
		Summary:      "Sign out",
		Description:  "Signing out will delete the current user session, so refreshing the token will no longer be possible.",
		Tags:         []string{"Auth"},
		Authenticate: true,
	})

	// Session
	handler.Add(Server, http.MethodGet, "/users/current/sessions", session.GetCurrentUserSessions, handler.Config{
		Summary:      "Get current user's sessions",
		Description:  "Get current user's sessions",
		Tags:         []string{"Session"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/users/current/sessions/current", session.GetCurrentUserCurrentSession, handler.Config{
		Summary:      "Get current user's active session",
		Description:  "Get current user's active session",
		Tags:         []string{"Session"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/users/current/sessions/:sessionId", session.GetCurrentUserSession, handler.Config{
		Summary:      "Get current user's session",
		Description:  "Get current user's session",
		Tags:         []string{"Session"},
		Authenticate: true,
	})

	// User
	handler.Add(Server, http.MethodGet, "/users/current", user.GetCurrentUser, handler.Config{
		Summary:      "Get current user",
		Description:  "Get current user",
		Tags:         []string{"User"},
		Authenticate: true,
	})

	// Verification
	handler.Add(Server, http.MethodPost, "/verify", verification.Verify, handler.Config{
		Summary:     "Verify user",
		Description: "Verify user using data taken from verification email",
		Tags:        []string{"Verification"},
	})
}
