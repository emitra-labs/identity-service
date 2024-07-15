package session_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/rest"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestGetCurrentUserCurrentSession_Success(t *testing.T) {
	session := tests.Data.Sessions[2]

	testkit.New(rest.Server).
		Get("/users/current/sessions/current").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", session.ID)).
		End()
}
