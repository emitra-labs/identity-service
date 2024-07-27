package user_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestUpdateUser_Success(t *testing.T) {
	data := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(tests.RESTServer).
		Patch("/users/"+tests.Data.Users[4].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(data).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.name", data["name"])).
		End()
}
