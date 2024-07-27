package auth_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestSignOut_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/auth/sign-out").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[3]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "User sign-out was successful")).
		End()
}
