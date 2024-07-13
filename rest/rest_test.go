package rest_test

import (
	"context"
	"os"
	"testing"

	"github.com/go-faker/faker/v4"
	dt "github.com/ukasyah-dev/common/db/testkit"
	"github.com/ukasyah-dev/common/mail"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func TestMain(m *testing.M) {
	dt.CreateTestDB()
	db.Open()
	mail.Open(os.Getenv("SMTP_URL"))
	setupTestData()
	code := m.Run()
	mail.Close()
	db.Close()
	dt.DestroyTestDB()
	os.Exit(code)
}

var testData struct {
	users []*model.User
}

func setupTestData() {
	ctx := context.Background()

	for i := 1; i <= 5; i++ {
		user, _ := controller.CreateUser(ctx, &model.CreateUserRequest{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "SuperSecret",
			Status:   constant.UserStatusActive,
		})
		testData.users = append(testData.users, user)
	}
}
