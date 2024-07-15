package session

import (
	"context"

	commonConstant "github.com/ukasyah-dev/common/constant"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/identity-service/model"
)

func GetCurrentUserCurrentSession(ctx context.Context, req *commonModel.Empty) (*model.Session, error) {
	sessionID, _ := ctx.Value(commonConstant.SessionID).(string)
	userID, _ := ctx.Value(commonConstant.UserID).(string)

	return GetSession(ctx, &model.GetSessionRequest{
		ID:     sessionID,
		UserID: userID,
	})
}
