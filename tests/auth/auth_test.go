package auth_test

import (
	"os"
	"testing"

	"github.com/ukasyah-dev/common/amqp"
	dt "github.com/ukasyah-dev/common/db/testkit"
	"github.com/ukasyah-dev/common/mail"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/tests"
)

func TestMain(m *testing.M) {
	amqp.Open(os.Getenv("AMQP_URL"))
	dt.CreateTestDB()
	db.Open()
	mail.Open(os.Getenv("SMTP_URL"))
	tests.Setup()

	code := m.Run()

	amqp.Close()
	mail.Close()
	db.Close()
	dt.DestroyTestDB()

	os.Exit(code)
}
