package verification_test

import (
	"os"
	"testing"

	"github.com/emitra-labs/identity-service/tests"
)

func TestMain(m *testing.M) {
	tests.Setup()
	defer tests.Teardown()
	os.Exit(m.Run())
}
