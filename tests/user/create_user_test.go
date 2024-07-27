package user_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestCreateUser_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/users").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    faker.Email(),
			"password": faker.Password(),
			"status":   "active",
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		End()
}
