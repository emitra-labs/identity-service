package verification

import (
	"context"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/id"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/constant"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
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
