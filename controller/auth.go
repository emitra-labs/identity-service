package controller

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func SignUp(ctx context.Context, req *model.SignUpRequest) (*commonModel.BasicResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	user := &model.User{
		ID:       id.New(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hash.Generate(req.Password),
		Status:   constant.UserStatusPendingVerification,
	}

	if err := db.DB.Create(user).Error; err != nil {
		log.Errorf("Failed to create user: %s", err)
		return nil, errors.Internal()
	}

	// TODO: Send verification email

	return &commonModel.BasicResponse{
		Message: "Check your email for further guidance.",
	}, nil
}
