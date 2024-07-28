package tests

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestGetCurrentUserCurrentSession_Success(t *testing.T) {
	session := data.sessions[2]

	testkit.New(restServer).
		Get("/users/current/sessions/current").
		Header("Authorization", "Bearer "+data.accessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", session.ID)).
		End()
}

func TestGetCurrentUserSession_Success(t *testing.T) {
	session := data.sessions[2]

	testkit.New(restServer).
		Get("/users/current/sessions/"+session.ID).
		Header("Authorization", "Bearer "+data.accessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", session.ID)).
		End()
}

func TestGetCurrentUserSessions_Success(t *testing.T) {
	testkit.New(restServer).
		Get("/users/current/sessions").
		Header("Authorization", "Bearer "+data.accessTokens[1]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		End()
}
