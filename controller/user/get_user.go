package user

import (
	"context"
	e "errors"

	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/log"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
	"gorm.io/gorm"
)

func GetUser(ctx context.Context, req *model.GetUserRequest) (*model.User, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx)

	if req.ID != "" {
		tx = tx.Where("id = ?", req.ID)
	}

	if req.Email != "" {
		tx = tx.Where("email = ?", req.Email)
	}

	if req.Status != "" {
		tx = tx.Where("status = ?", req.Status)
	}

	var result model.User

	if err := tx.Take(&result).Error; err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get user: %s", err)
		return nil, errors.Internal()
	}

	return &result, nil
}
