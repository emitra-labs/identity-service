package session_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestGetCurrentUserCurrentSession_Success(t *testing.T) {
	session := tests.Data.Sessions[2]

	testkit.New(tests.RESTServer).
		Get("/users/current/sessions/current").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", session.ID)).
		End()
}
