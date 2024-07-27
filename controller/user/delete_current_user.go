package user

import (
	"context"

	"github.com/emitra-labs/common/constant"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/identity-service/model"
)

func DeleteCurrentUser(ctx context.Context, req *commonModel.Empty) (*model.User, error) {
	userID, _ := ctx.Value(constant.UserID).(string)
	return DeleteUser(ctx, &model.DeleteUserRequest{ID: userID})
}
