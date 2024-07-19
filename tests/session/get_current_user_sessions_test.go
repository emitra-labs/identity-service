package session_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestGetCurrentUserSessions_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/users/current/sessions").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[1]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		End()
}
