package user_test

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/emitra-labs/identity-service/tests"
)

func TestDeleteCurrentUser_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Delete("/users/current").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
