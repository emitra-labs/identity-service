package tests

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/emitra-labs/common/amqp"
	commonAuth "github.com/emitra-labs/common/auth"
	dt "github.com/emitra-labs/common/db/testkit"
	"github.com/emitra-labs/identity-service/constant"
	"github.com/emitra-labs/identity-service/controller"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
	"github.com/emitra-labs/identity-service/rest"
	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
)

var data struct {
	accessTokens  []string
	sessions      []*model.Session
	users         []*model.User
	verifications []*model.Verification
}

var restServer = rest.NewServer()

func TestMain(m *testing.M) {
	setup()
	defer teardown()
	os.Exit(m.Run())
}

func setup() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	dt.CreateTestDB()
	db.Open()

	ctx := context.Background()

	jwtPrivateKey, err := jwt.ParseEdPrivateKeyFromPEM([]byte(os.Getenv("JWT_PRIVATE_KEY")))
	if err != nil {
		panic(err)
	}

	for i := 0; i <= 4; i++ {
		status := constant.UserStatusActive

		if i == 4 {
			status = constant.UserStatusPendingVerification
		}

		user, _ := controller.CreateUser(ctx, &model.CreateUserRequest{
			Name:       faker.Name(),
			Email:      faker.Email(),
			Password:   "SuperSecret",
			Status:     status,
			SuperAdmin: i == 0,
		})
		data.users = append(data.users, user)

		if status == constant.UserStatusActive {
			session, _ := controller.CreateSession(ctx, &model.CreateSessionRequest{
				UserID:    user.ID,
				ExpiresAt: time.Now().AddDate(0, 0, 1),
			})
			data.sessions = append(data.sessions, session)

			accessToken, _ := commonAuth.GenerateAccessToken(jwtPrivateKey, commonAuth.Claims{
				SessionID:  session.ID,
				SuperAdmin: user.SuperAdmin,
				UserID:     user.ID,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
				},
			})
			data.accessTokens = append(data.accessTokens, accessToken)

		} else if status == constant.UserStatusPendingVerification {
			verification, _ := controller.CreateVerification(ctx, &model.CreateVerificationRequest{
				UserID:    user.ID,
				ExpiresAt: time.Now().Add(15 * time.Minute),
			})
			data.verifications = append(data.verifications, verification)
		}
	}
}

func teardown() {
	amqp.Close()
	db.Close()
	dt.DestroyTestDB()
}
