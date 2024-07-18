package verification

import (
	"context"
	e "errors"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	commonModel "github.com/ukasyah-dev/common/model"
	"github.com/ukasyah-dev/identity-service/constant"
	"github.com/ukasyah-dev/identity-service/controller/user"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
	"gorm.io/gorm"
)

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

	// Set user status to active
	_, err = user.UpdateUser(ctx, &model.UpdateUserRequest{
		ID:     verification.UserID,
		Status: constant.UserStatusActive,
	})
	if err != nil {
		return nil, err
	}

	// Set verification status to success
	verification.Status = constant.VerificationStatusSuccess
	if err := db.DB.Save(&verification).Error; err != nil {
		log.Errorf("Failed to update verification status: %s", err)
		return nil, err
	}

	return &commonModel.BasicResponse{
		Message: "Your account is verified successfully.",
	}, nil
}
