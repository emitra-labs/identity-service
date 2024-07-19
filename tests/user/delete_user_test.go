package user_test

import (
	"net/http"
	"testing"

	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestDeleteUser_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Delete("/users/"+tests.Data.Users[3].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
