package rest

import (
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/server"
	"github.com/ukasyah-dev/identity-service/controller"
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
	spec.SetHTTPBearerTokenSecurity("BearerAuth", "JWT", "")

	Server = server.New(server.Config{
		OpenAPI: server.OpenAPI{Spec: &spec},
	})

	handler.Add(Server, http.MethodGet, "/", controller.HealthCheck, handler.Config{
		Summary:     "Health check",
		Description: "Check whether the server is ready to serve.",
		Tags:        []string{"Health"},
	})

	// Auth
	handler.Add(Server, http.MethodPost, "/auth/sign-up", controller.SignUp, handler.Config{
		Summary:     "Sign up",
		Description: "Sign up for a new user account.",
		Tags:        []string{"Auth"},
	})

	handler.Add(Server, http.MethodGet, "/hello/:name", controller.Hello, handler.Config{
		Summary:     "Say hello",
		Description: "We will greet you with warmth and kindness.",
		Tags:        []string{"Hello"},
	})
}
