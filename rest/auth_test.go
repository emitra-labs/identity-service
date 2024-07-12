package rest_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/rest"
)

func TestSignUp_Success(t *testing.T) {
	testkit.New(rest.Server).Debug().
		Post("/auth/sign-up").
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    faker.Email(),
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "TODO: Sign up")).
		End()
}
