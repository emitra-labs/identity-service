package session

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
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
