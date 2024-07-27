package auth_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestRefreshToken_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/refresh").
		JSON(map[string]any{
			"token":  tests.Data.Sessions[1].Token,
			"userId": tests.Data.Users[1].ID,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.access.token")).
		Assert(jsonpath.NotEqual("$.refresh.token", tests.Data.Sessions[1].Token)).
		Assert(jsonpath.Equal("$.user.id", tests.Data.Users[1].ID)).
		End()
}
