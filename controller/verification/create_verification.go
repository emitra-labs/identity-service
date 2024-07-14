package verification

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func CreateVerification(ctx context.Context, req *model.CreateVerificationRequest) (*model.Verification, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	verification := &model.Verification{
		ID:        id.New(),
		Token:     id.New(36),
		UserID:    req.UserID,
		Status:    constant.VerificationStatusPending,
		ExpiresAt: req.ExpiresAt,
	}

	if err := db.DB.WithContext(ctx).Create(verification).Error; err != nil {
		log.Errorf("Failed to create verification: %s", err)
		return nil, errors.Internal()
	}

	return verification, nil
}
