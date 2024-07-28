package tests

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestVerify_Success(t *testing.T) {
	verification := data.verifications[0]

	testkit.New(restServer).
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
