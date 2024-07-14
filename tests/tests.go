package tests

import (
	"context"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	commonAuth "github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/controller/verification"
	"github.com/ukasyah-dev/identity-service/model"
)

var Data struct {
	AccessTokens  []string
	Users         []*model.User
	Verifications []*model.Verification
}

func Setup() {
	ctx := context.Background()

	jwtPrivateKey, err := commonAuth.ParsePrivateKeyFromBase64(os.Getenv("BASE64_JWT_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

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

		if status == constant.UserStatusActive {
			accessToken, _ := commonAuth.GenerateAccessToken(jwtPrivateKey, commonAuth.Claims{
				UserID: u.ID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
				},
			})

			Data.AccessTokens = append(Data.AccessTokens, accessToken)
		} else if status == constant.UserStatusPendingVerification {
			verification, _ := verification.CreateVerification(ctx, &model.CreateVerificationRequest{
				UserID:    u.ID,
				ExpiresAt: time.Now().Add(15 * time.Minute),
			})

			Data.Verifications = append(Data.Verifications, verification)
		}
	}
}
