package tests

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestSignUp_Success(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/sign-up").
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    faker.Email(),
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Check your email")).
		End()
}

func TestSignUp_EmailExists(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/sign-up").
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    data.users[0].Email,
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "Check your email")).
		End()
}

func TestSignIn_Success(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    data.users[0].Email,
			"password": "SuperSecret",
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.access.token")).
		End()
}

func TestSignIn_InvalidEmailOrPassword(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    faker.Email(),
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusBadRequest).
		Assert(jsonpath.Contains("$.error", "Invalid email or password")).
		End()

	testkit.New(restServer).
		Post("/auth/sign-in").
		JSON(map[string]any{
			"email":    data.users[1].Email,
			"password": faker.Password(),
		}).
		Expect(t).
		Status(http.StatusBadRequest).
		Assert(jsonpath.Contains("$.error", "Invalid email or password")).
		End()
}

func TestRefreshToken_Success(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/refresh").
		JSON(map[string]any{
			"token":  data.sessions[1].Token,
			"userId": data.users[1].ID,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.access.token")).
		Assert(jsonpath.NotEqual("$.refresh.token", data.sessions[1].Token)).
		Assert(jsonpath.Equal("$.user.id", data.users[1].ID)).
		End()
}

func TestSignOut_Success(t *testing.T) {
	testkit.New(restServer).
		Post("/auth/sign-out").
		Header("Authorization", "Bearer "+data.accessTokens[3]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Contains("$.message", "User sign-out was successful")).
		End()
}
