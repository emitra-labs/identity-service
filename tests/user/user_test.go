package user_test

import (
	"os"
	"testing"

	"github.com/ukasyah-dev/identity-service/tests"
)

func TestMain(m *testing.M) {
	tests.Setup()
	defer tests.Teardown()
	os.Exit(m.Run())
}
