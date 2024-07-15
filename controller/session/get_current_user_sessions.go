package session

import (
	"context"

	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/identity-service/model"
)

func GetCurrentUserSessions(ctx context.Context, req *model.GetCurrentUserSessionsRequest) (*model.GetSessionsResponse, error) {
	userID, _ := ctx.Value(constant.UserID).(string)

	return GetSessions(ctx, &model.GetSessionsRequest{
		PaginationRequest: req.PaginationRequest,
		UserID:            userID,
	})
}
