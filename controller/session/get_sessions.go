package session

import (
	"context"

	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func GetSessions(ctx context.Context, req *model.GetSessionsRequest) (*model.GetSessionsResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).
		Model(&model.Session{}).
		Where("user_id = ?", req.UserID)

	data, pagination, err := paginator.Paginate[model.Session](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetSessionsResponse{
		Data:               data,
		PaginationResponse: *pagination,
	}, nil
}
