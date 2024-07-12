package controller

import (
	"context"

	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/model"
)

func SignUp(ctx context.Context, req *model.SignUpRequest) (*commonModel.BasicResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	return &commonModel.BasicResponse{Message: "TODO: Sign up"}, nil
}
