package controller

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/emitra-labs/common/auth"
	commonConstant "github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/hash"
	"github.com/emitra-labs/common/id"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/mail"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/constant"
	"github.com/emitra-labs/identity-service/controller/session"
	"github.com/emitra-labs/identity-service/controller/user"
	"github.com/emitra-labs/identity-service/controller/verification"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
	"github.com/golang-jwt/jwt/v5"
)

var accessExpiresIn = 30 * time.Minute

func SignUp(ctx context.Context, req *model.SignUpRequest) (*commonModel.BasicResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	alreadyExists := false

	u, err := user.CreateUser(ctx, &model.CreateUserRequest{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Status:   constant.UserStatusPendingVerification,
	})
	if err != nil {
		if errors.IsAlreadyExists(err) {
			alreadyExists = true
		} else {
			return nil, err
		}
	}

	if alreadyExists {
		err = mail.Send(mail.Email{
			From:    os.Getenv("EMAIL_FROM"),
			To:      req.Email,
			Subject: "Sign in to Your Account",
			Body: mail.Body{
				Name:   req.Name,
				Intros: []string{"Your email has been registered. Please click the link below to proceed."},
				Actions: []mail.Action{
					{
						Text: "Sign in",
						Link: os.Getenv("EMAIL_AUTH_URL") + "/sign-in",
					},
				},
			},
		})
		if err != nil {
			log.Errorf("Failed to send sign-in email: %s", err)
			return nil, errors.Internal()
		}
	} else {
		v, err := verification.CreateVerification(ctx, &model.CreateVerificationRequest{
			UserID:    u.ID,
			ExpiresAt: time.Now().Add(15 * time.Minute),
		})
		if err != nil {
			return nil, err
		}

		err = mail.Send(mail.Email{
			From:    os.Getenv("EMAIL_FROM"),
			To:      req.Email,
			Subject: "Verify Your Account",
			Body: mail.Body{
				Name:   req.Name,
				Intros: []string{"You need to pass the verification step. Please click the link below to proceed."},
				Actions: []mail.Action{
					{
						Text: "Verify your account",
						Link: fmt.Sprintf("%s/verify?userId=%s&token=%s", os.Getenv("EMAIL_AUTH_URL"), u.ID, v.Token),
					},
				},
			},
		})
		if err != nil {
			log.Errorf("Failed to send verification email: %s", err)
			return nil, errors.Internal()
		}
	}

	return &commonModel.BasicResponse{
		Message: "Check your email for further guidance.",
	}, nil
}

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

	accessExpiresAt := time.Now().Add(accessExpiresIn)

	accessToken, err := auth.GenerateAccessToken(jwtPrivateKey, auth.Claims{
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
		Access: &model.Access{
			Token:     accessToken,
			ExpiresAt: accessExpiresAt,
		},
		Refresh: &model.Refresh{
			Token:     s.Token,
			ExpiresAt: s.ExpiresAt,
		},
		User: u,
	}, nil
}

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
	accessToken, err := auth.GenerateAccessToken(jwtPrivateKey, auth.Claims{
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

func SignOut(ctx context.Context, req *commonModel.Empty) (*commonModel.BasicResponse, error) {
	sessionID, _ := ctx.Value(commonConstant.SessionID).(string)

	_, err := session.DeleteSession(ctx, &model.DeleteSessionRequest{
		ID: sessionID,
	})
	if err != nil {
		return nil, err
	}

	return &commonModel.BasicResponse{
		Message: "User sign-out was successful",
	}, nil
}
