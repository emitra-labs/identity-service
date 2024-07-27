package user_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestGetUser_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/users/"+tests.Data.Users[1].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		Assert(jsonpath.Equal("$.email", tests.Data.Users[1].Email)).
		End()
}
