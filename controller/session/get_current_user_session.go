package session

import (
	"context"

	"github.com/ukasyah-dev/identity-service/model"
)

func GetCurrentUserSession(ctx context.Context, req *model.GetCurrentUserSession) (*model.Session, error) {
	return GetSession(ctx, &model.GetSessionRequest{
		ID: req.ID,
	})
}
