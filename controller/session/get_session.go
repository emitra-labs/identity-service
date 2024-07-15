package session

import (
	"context"
	e "errors"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
	"gorm.io/gorm"
)

func GetSession(ctx context.Context, req *model.GetSessionRequest) (*model.Session, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	session := new(model.Session)

	tx := db.DB.WithContext(ctx).Where("expires_at > NOW()")

	if req.ID != "" {
		tx = tx.Where("id = ?", req.ID)
	}

	if req.Token != "" {
		tx = tx.Where("token = ?", req.Token)
	}

	if req.UserID != "" {
		tx = tx.Where("user_id = ?", req.UserID)
	}

	err := tx.Take(session).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("Session not found")
		}

		log.Errorf("Failed to view session: %s", err)
		return nil, errors.Internal()
	}

	return session, nil
}
