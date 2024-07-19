package user_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
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
