package controller

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
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
		log.Debug("TODO: Ask user to login via email")
	} else {
		log.Debug("TODO: Send verification link via email")
	}

	return &commonModel.BasicResponse{
		Message: "Check your email for further guidance.",
	}, nil
}
