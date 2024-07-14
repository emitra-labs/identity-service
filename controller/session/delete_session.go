package session

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func DeleteOldSessions(ctx context.Context, req *model.DeleteOldSessionsRequest) (*model.Empty, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Delete old sessions, preserve 3 newest
	sql := `DELETE FROM sessions WHERE user_id = ? AND id NOT IN (
		SELECT id FROM sessions WHERE user_id = ? ORDER BY created_at DESC LIMIT 3
	)`
	if err := db.DB.Exec(sql, req.UserID, req.UserID).Error; err != nil {
		log.Errorf("Failed to delete old sessions: %s", err)
		return nil, errors.Internal()
	}

	return &model.Empty{}, nil
}
