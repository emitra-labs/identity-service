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
		Description: "Check whether the server is ready to serve.",
		Tags:        []string{"Health"},
	})

	// Auth
	handler.Add(Server, http.MethodPost, "/auth/sign-up", auth.SignUp, handler.Config{
		Summary:     "Sign up",
		Description: "Sign up for a new user account.",
		Tags:        []string{"Auth"},
	})
	handler.Add(Server, http.MethodPost, "/auth/sign-in", auth.SignIn, handler.Config{
		Summary:     "Sign in",
		Description: "Sign in.",
		Tags:        []string{"Auth"},
	})
	handler.Add(Server, http.MethodPost, "/auth/refresh", auth.RefreshToken, handler.Config{
		Summary:     "Refresh token",
		Description: "Generate new access token.",
		Tags:        []string{"Auth"},
	})

	// User
	handler.Add(Server, http.MethodGet, "/users/current", user.GetCurrentUser, handler.Config{
		Summary:      "Get current user",
		Description:  "Get current user.",
		Tags:         []string{"User"},
		Authenticate: true,
	})

	// Verification
	handler.Add(Server, http.MethodPost, "/verify", verification.Verify, handler.Config{
		Summary:     "Verify user",
		Description: "Verify user using data taken from verification email.",
		Tags:        []string{"Verification"},
	})
}
