package session_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
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
