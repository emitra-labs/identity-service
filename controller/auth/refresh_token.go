package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/identity-service/controller/session"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func RefreshToken(ctx context.Context, req *model.RefreshTokenRequest) (*model.SignInResponse, error) {
	session, err := session.GetSession(ctx, &model.GetSessionRequest{
		Token:  req.Token,
		UserID: req.UserID,
	})
	if err != nil {
		return nil, err
	}

	u, err := user.GetUser(ctx, &model.GetUserRequest{
		ID: session.UserID,
	})
	if err != nil {
		return nil, err
	}

	accessExpiresAt := time.Now().Add(accessExpiresIn)

	// Generate new access token
	accessToken, err := auth.GenerateAccessToken(privateKey, auth.Claims{
		UserID:     u.ID,
		SessionID:  session.ID,
		SuperAdmin: u.SuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiresAt),
		},
	})
	if err != nil {
		return nil, err
	}

	// Generate new refresh token
	session.Token = id.New(48)
	if err := db.DB.Save(session).Error; err != nil {
		log.Errorf("Failed to save session: %s", err)
		return nil, errors.Internal()
	}

	return &model.SignInResponse{
		Access: &model.Access{
			Token:     accessToken,
			ExpiresAt: accessExpiresAt,
		},
		Refresh: &model.Refresh{
			Token:     session.Token,
			ExpiresAt: session.ExpiresAt,
		},
		User: u,
	}, nil
}
