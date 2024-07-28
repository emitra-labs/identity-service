package tests

import (
	"net/http"
	"testing"

	"github.com/emitra-labs/common/rest/testkit"
	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
)

func TestCreateUser_Success(t *testing.T) {
	testkit.New(restServer).
		Post("/users").
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		JSON(map[string]any{
			"name":     faker.Name(),
			"email":    faker.Email(),
			"password": faker.Password(),
			"status":   "active",
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		End()
}

func TestGetUsers_Success(t *testing.T) {
	testkit.New(restServer).
		Get("/users").
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		End()
}

func TestGetUsers_PermissionDenied(t *testing.T) {
	testkit.New(restServer).
		Get("/users").
		Header("Authorization", "Bearer "+data.accessTokens[1]).
		Expect(t).
		Status(http.StatusForbidden).
		End()
}

func TestGetUser_Success(t *testing.T) {
	testkit.New(restServer).
		Get("/users/"+data.users[1].ID).
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		Assert(jsonpath.Equal("$.email", data.users[1].Email)).
		End()
}

func TestGetCurrentUser_Success(t *testing.T) {
	testkit.New(restServer).
		Get("/users/current").
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		Assert(jsonpath.Equal("$.email", data.users[0].Email)).
		End()
}

func TestUpdateUser_Success(t *testing.T) {
	user := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(restServer).
		Patch("/users/"+data.users[4].ID).
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		JSON(user).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.name", user["name"])).
		End()
}

func TestUpdateCurrentUser_Success(t *testing.T) {
	user := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(restServer).
		Patch("/users/current").
		Header("Authorization", "Bearer "+data.accessTokens[1]).
		JSON(user).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.name", user["name"])).
		End()
}

func TestDeleteUser_Success(t *testing.T) {
	testkit.New(restServer).
		Delete("/users/"+data.users[3].ID).
		Header("Authorization", "Bearer "+data.accessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func TestDeleteCurrentUser_Success(t *testing.T) {
	testkit.New(restServer).
		Delete("/users/current").
		Header("Authorization", "Bearer "+data.accessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
