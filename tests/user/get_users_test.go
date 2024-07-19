package user_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestGetUsers_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/users").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		End()
}

func TestGetUsers_PermissionDenied(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/users").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[1]).
		Expect(t).
		Status(http.StatusForbidden).
		End()
}
