package user

import (
	"context"
	"os"

	"github.com/emitra-labs/common/amqp"
	commonConstant "github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/hash"
	"github.com/emitra-labs/common/log"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
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

	if os.Getenv("SKIP_AMQP_PUBLISHING") != "true" {
		err := amqp.Publish(ctx, "user-mutation", &commonModel.Mutation[model.User]{
			Type: commonConstant.MutationUpdated,
			Data: user,
		})
		if err != nil {
			log.Errorf("Failed to publish to user-mutation: %s", err)
			return nil, errors.Internal()
		}
	}

	return user, nil
}
