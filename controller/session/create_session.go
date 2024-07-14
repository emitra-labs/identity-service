package session

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.Session, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	session := &model.Session{
		ID:        id.New(),
		Token:     id.New(36),
		UserID:    req.UserID,
		ExpiresAt: req.ExpiresAt,
	}

	if err := db.DB.WithContext(ctx).Create(session).Error; err != nil {
		log.Errorf("Failed to create session: %s", err)
		return nil, errors.Internal()
	}

	return session, nil
}
