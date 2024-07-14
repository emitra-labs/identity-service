package rest_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-faker/faker/v4"
	dt "github.com/ukasyah-dev/common/db/testkit"
	"github.com/ukasyah-dev/common/mail"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/controller/verification"
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
	users         []*model.User
	verifications []*model.Verification
}

func setupTestData() {
	ctx := context.Background()

	for i := 0; i <= 4; i++ {
		status := constant.UserStatusActive

		if i == 4 {
			status = constant.UserStatusPendingVerification
		}

		u, _ := user.CreateUser(ctx, &model.CreateUserRequest{
			Name:     faker.Name(),
			Email:    faker.Email(),
			Password: "SuperSecret",
			Status:   status,
		})

		testData.users = append(testData.users, u)

		if status == constant.UserStatusPendingVerification {
			verification, _ := verification.CreateVerification(ctx, &model.CreateVerificationRequest{
				UserID:    u.ID,
				ExpiresAt: time.Now().Add(15 * time.Minute),
			})

			testData.verifications = append(testData.verifications, verification)
		}
	}
}
