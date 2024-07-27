package session

import (
	"context"

	"github.com/emitra-labs/common/paginator"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
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
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
