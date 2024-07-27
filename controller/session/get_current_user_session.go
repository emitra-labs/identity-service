package session

import (
	"context"

	"github.com/emitra-labs/identity-service/model"
)

func GetCurrentUserSession(ctx context.Context, req *model.GetCurrentUserSession) (*model.Session, error) {
	return GetSession(ctx, &model.GetSessionRequest{
		ID: req.ID,
	})
}
