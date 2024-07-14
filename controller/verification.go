package controller

import (
	"context"
	e "errors"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
	"gorm.io/gorm"
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

func Verify(ctx context.Context, req *model.VerifyRequest) (*commonModel.BasicResponse, error) {
	var verification model.Verification

	err := db.DB.WithContext(ctx).Preload("User").
		Where("user_id = ?", req.UserID).
		Where("token = ?", req.Token).
		Where("status = ?", constant.VerificationStatusPending).
		Where("expires_at > NOW()").
		Take(&verification).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.InvalidArgument("Failed to verify using provided token. It may be expired or invalid.")
		}

		log.Errorf("Failed to verify: %s", err)
		return nil, errors.Internal()
	}

	verification.User.Status = constant.UserStatusActive

	if err := db.DB.Save(verification.User).Error; err != nil {
		log.Errorf("Failed to update user status: %s", err)
		return nil, err
	}

	return &commonModel.BasicResponse{
		Message: "Your account has been verified.",
	}, nil
}
