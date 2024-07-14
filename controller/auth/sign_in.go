package auth

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/session"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/model"
)

func SignIn(ctx context.Context, req *model.SignInRequest) (*model.SignInResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	u, err := user.GetUser(ctx, &model.GetUserRequest{
		Email:  req.Email,
		Status: constant.UserStatusActive,
	})
	if err != nil {
		if errors.IsNotFound(err) {
			return nil, errors.InvalidArgument("Invalid email or password")
		}
		return nil, err
	}

	if err := hash.Verify(req.Password, u.Password); err != nil {
		if errors.IsInvalidArgument(err) {
			return nil, errors.InvalidArgument("Invalid email or password")
		}
		return nil, err
	}

	s, err := session.CreateSession(ctx, &model.CreateSessionRequest{
		UserID:    u.ID,
		ExpiresAt: time.Now().AddDate(0, 0, 7),
	})
	if err != nil {
		return nil, err
	}

	accessExpiresAt := time.Now().Add(30 * time.Minute)

	accessToken, err := auth.GenerateAccessToken(privateKey, auth.Claims{
		UserID:     u.ID,
		SessionID:  s.ID,
		SuperAdmin: u.SuperAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiresAt),
		},
	})
	if err != nil {
		return nil, err
	}

	_, err = session.DeleteOldSessions(ctx, &model.DeleteOldSessionsRequest{
		UserID: u.ID,
	})
	if err != nil {
		return nil, err
	}

	return &model.SignInResponse{
		Access: model.Access{
			Token:     accessToken,
			ExpiresAt: accessExpiresAt,
		},
		Refresh: model.Refresh{
			Token:     s.Token,
			ExpiresAt: s.ExpiresAt,
		},
	}, nil
}
