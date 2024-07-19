package auth_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestSignUp_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/sign-up").
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    faker.Email(),
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Check your email")).
		End()
}

func TestSignUp_EmailExists(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/sign-up").
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    tests.Data.Users[0].Email,
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Check your email")).
		End()
}
