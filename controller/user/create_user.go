package user

import (
	"context"
	e "errors"
	"os"

	"github.com/ukasyah-dev/common/amqp"
	commonConstant "github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	commonModel "github.com/ukasyah-dev/common/model"
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

	if os.Getenv("SKIP_AMQP_PUBLISHING") != "true" {
		err := amqp.Publish(ctx, "user-mutation", &commonModel.Mutation[model.User]{
			Type: commonConstant.MutationCreated,
			Data: user,
		})
		if err != nil {
			log.Errorf("Failed to publish to user-mutation: %s", err)
			return nil, errors.Internal()
		}
	}

	return user, nil
}
