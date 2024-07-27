package user

import (
	"context"

	"github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/identity-service/model"
)

func UpdateCurrentUser(ctx context.Context, req *model.UpdateCurrentUserRequest) (*model.User, error) {
	userID, _ := ctx.Value(constant.UserID).(string)
	return UpdateUser(ctx, &model.UpdateUserRequest{
		ID:       userID,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	})
}
