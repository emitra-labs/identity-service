package auth_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
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
