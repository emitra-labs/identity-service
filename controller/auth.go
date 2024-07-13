package controller

import (
	"context"
	"os"

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

	_, err := CreateUser(ctx, &model.CreateUserRequest{
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
	} else {
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
						Link: os.Getenv("EMAIL_AUTH_URL") + "/verify",
					},
				},
			},
		})
	}

	if err != nil {
		log.Errorf("Failed to send email: %s", err)
		return nil, errors.Internal()
	}

	return &commonModel.BasicResponse{
		Message: "Check your email for further guidance.",
	}, nil
}
