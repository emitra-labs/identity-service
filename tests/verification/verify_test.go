package verification_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/rest"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestVerify_Success(t *testing.T) {
	verification := tests.Data.Verifications[0]

	testkit.New(rest.Server).
		Post("/verify").
		JSON(map[string]any{
			"userId": verification.UserID,
			"token":  verification.Token,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Your account has been verified")).
		End()
}
