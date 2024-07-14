package user

import (
	"context"
	e "errors"

	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"github.com/ukasyah-dev/identity-service/db"
	"github.com/ukasyah-dev/identity-service/model"
	"gorm.io/gorm"
)

func CreateUser(ctx context.Context, req *model.CreateUserRequest) (*model.User, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	user := &model.User{
		ID:         id.New(),
		Name:       req.Name,
		Email:      req.Email,
		Password:   hash.Generate(req.Password),
		Status:     req.Status,
		SuperAdmin: req.SuperAdmin,
	}

	if err := db.DB.WithContext(ctx).Create(user).Error; err != nil {
		if e.Is(err, gorm.ErrDuplicatedKey) {
			return nil, errors.AlreadyExists()
		}

		log.Errorf("Failed to create user: %s", err)
		return nil, errors.Internal()
	}

	return user, nil
}
