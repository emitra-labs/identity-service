package session

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
)

func DeleteOldSessions(ctx context.Context, req *model.DeleteOldSessionsRequest) (*commonModel.Empty, error) {
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

	return &commonModel.Empty{}, nil
}
