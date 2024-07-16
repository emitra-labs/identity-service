package user

import (
	"context"

	"github.com/ukasyah-dev/common/constant"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/identity-service/model"
)

func DeleteCurrentUser(ctx context.Context, req *commonModel.Empty) (*model.User, error) {
	userID, _ := ctx.Value(constant.UserID).(string)
	return DeleteUser(ctx, &model.DeleteUserRequest{ID: userID})
}
