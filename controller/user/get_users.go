package user

import (
	"context"

	"github.com/emitra-labs/common/paginator"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
)

func GetUsers(ctx context.Context, req *model.GetUsersRequest) (*model.GetUsersResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).Model(&model.User{})

	data, pagination, err := paginator.Paginate[model.User](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetUsersResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
