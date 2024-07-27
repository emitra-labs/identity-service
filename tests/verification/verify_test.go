package verification_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestVerify_Success(t *testing.T) {
	verification := tests.Data.Verifications[0]

	testkit.New(tests.RESTServer).
		Post("/verify").
		JSON(map[string]any{
			"userId": verification.UserID,
			"token":  verification.Token,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Your account is verified successfully.")).
		End()
}
