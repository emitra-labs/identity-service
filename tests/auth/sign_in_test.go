package auth_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestSignIn_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    tests.Data.Users[0].Email,
			"password": "SuperSecret",
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.access.token")).
		End()
}

func TestSignIn_InvalidEmailOrPassword(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    faker.Email(),
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusBadRequest).
		Assert(jsonpath.Contains("$.error", "Invalid email or password")).
		End()

	testkit.New(tests.RESTServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    tests.Data.Users[1].Email,
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusBadRequest).
		Assert(jsonpath.Contains("$.error", "Invalid email or password")).
		End()
}
