package rest_test

import (
	"os"
	"testing"

	dt "github.com/ukasyah-dev/common/db/testkit"
	"github.com/ukasyah-dev/identity-service/db"
)

func TestMain(m *testing.M) {
	dt.CreateTestDB()
	db.Open()
	code := m.Run()
	db.Close()
	dt.DestroyTestDB()
	os.Exit(code)
}
