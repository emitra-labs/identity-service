package user_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestUpdateCurrentUser_Success(t *testing.T) {
	data := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(tests.RESTServer).
		Patch("/users/current").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[1]).
		JSON(data).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.name", data["name"])).
		End()
}
