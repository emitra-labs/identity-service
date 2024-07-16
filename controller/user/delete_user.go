package user

import (
	"context"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
)

func DeleteUser(ctx context.Context, req *model.DeleteUserRequest) (*model.User, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	user, err := GetUser(ctx, &model.GetUserRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	err = db.DB.WithContext(ctx).Delete(user).Error
	if err != nil {
		return nil, errors.Internal()
	}

	return user, nil
}
