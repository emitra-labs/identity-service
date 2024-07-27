package session

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
)

func DeleteSession(ctx context.Context, req *model.DeleteSessionRequest) (*model.Session, error) {
	s, err := GetSession(ctx, &model.GetSessionRequest{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := db.DB.Delete(s).Error; err != nil {
		log.Errorf("Failed to delete session: %s", err)
		return nil, errors.Internal()
	}

	return s, nil
}
