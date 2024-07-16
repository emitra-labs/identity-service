package user

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func UpdateUser(ctx context.Context, req *model.UpdateUserRequest) (*model.User, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	user, err := GetUser(ctx, &model.GetUserRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	if req.Password != "" {
		user.Password = hash.Generate(req.Password)
	}

	if req.Status != "" {
		user.Status = req.Status
	}

	err = db.DB.WithContext(ctx).Save(user).Error
	if err != nil {
		log.Errorf("Failed to update user: %s", err)
		return nil, errors.Internal()
	}

	return user, nil
}
