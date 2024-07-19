package tests

import (
	"context"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/common/amqp"
	commonAuth "github.com/ukasyah-dev/common/auth"
	dt "github.com/ukasyah-dev/common/db/testkit"
	"github.com/ukasyah-dev/common/mail"
	restServer "github.com/ukasyah-dev/common/rest/server"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/session"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/controller/verification"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
	"github.com/ukasyah-dev/identity-service/rest"
)

var Data struct {
	AccessTokens  []string
	Sessions      []*model.Session
	Users         []*model.User
	Verifications []*model.Verification
}

var RESTServer *restServer.Server

func Setup() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	dt.CreateTestDB()
	db.Open()
	mail.Open(os.Getenv("SMTP_URL"))
	RESTServer = rest.NewServer()

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
			Name:       faker.Name(),
			Email:      faker.Email(),
			Password:   "SuperSecret",
			Status:     status,
			SuperAdmin: i == 0,
		})
		Data.Users = append(Data.Users, u)

		if status == constant.UserStatusActive {
			s, _ := session.CreateSession(ctx, &model.CreateSessionRequest{
				UserID:    u.ID,
				ExpiresAt: time.Now().AddDate(0, 0, 1),
			})
			Data.Sessions = append(Data.Sessions, s)

			accessToken, _ := commonAuth.GenerateAccessToken(jwtPrivateKey, commonAuth.Claims{
				SessionID:  s.ID,
				SuperAdmin: u.SuperAdmin,
				UserID:     u.ID,
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

func Teardown() {
	amqp.Close()
	mail.Close()
	db.Close()
	dt.DestroyTestDB()
}
