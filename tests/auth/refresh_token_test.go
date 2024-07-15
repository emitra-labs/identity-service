package auth_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/rest"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestRefreshToken_Success(t *testing.T) {
	testkit.New(rest.Server).
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
