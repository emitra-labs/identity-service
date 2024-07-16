package user_test

import (
	"net/http"
	"testing"

	"github.com/ukasyah-dev/common/rest/testkit"
	"github.com/ukasyah-dev/identity-service/rest"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestDeleteCurrentUser_Success(t *testing.T) {
	testkit.New(rest.Server).
		Delete("/users/current").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
