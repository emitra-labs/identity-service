package tests

import (
	"context"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/controller/verification"
	"github.com/ukasyah-dev/identity-service/model"
)

var Data struct {
	Users         []*model.User
	Verifications []*model.Verification
}

func Setup() {
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

		Data.Users = append(Data.Users, u)

		if status == constant.UserStatusPendingVerification {
			verification, _ := verification.CreateVerification(ctx, &model.CreateVerificationRequest{
				UserID:    u.ID,
				ExpiresAt: time.Now().Add(15 * time.Minute),
			})

			Data.Verifications = append(Data.Verifications, verification)
		}
	}
}