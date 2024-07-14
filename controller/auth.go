package controller

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/mail"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/model"
)

func SignUp(ctx context.Context, req *model.SignUpRequest) (*commonModel.BasicResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	alreadyExists := false

	user, err := CreateUser(ctx, &model.CreateUserRequest{
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
		verification, err := CreateVerification(ctx, &model.CreateVerificationRequest{
			UserID:    user.ID,
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
						Link: fmt.Sprintf("%s/verify?userId=%s&token=%s", os.Getenv("EMAIL_AUTH_URL"), user.ID, verification.Token),
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
