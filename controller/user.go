package controller

import (
	"context"
	e "errors"
	"os"

	"github.com/emitra-labs/common/amqp"
	commonConstant "github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/hash"
	"github.com/emitra-labs/common/id"
	"github.com/emitra-labs/common/log"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/common/paginator"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
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

func GetUsers(ctx context.Context, req *model.GetUsersRequest) (*model.GetUsersResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).Model(&model.User{})

	data, pagination, err := paginator.Paginate[model.User](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetUsersResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}

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

func GetCurrentUser(ctx context.Context, req *commonModel.Empty) (*model.User, error) {
	userID, _ := ctx.Value(commonConstant.UserID).(string)
	return GetUser(ctx, &model.GetUserRequest{ID: userID})
}

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

	if os.Getenv("SKIP_AMQP_PUBLISHING") != "true" {
		err := amqp.Publish(ctx, "user-mutation", &commonModel.Mutation[model.User]{
			Type: commonConstant.MutationDeleted,
			Data: user,
		})
		if err != nil {
			log.Errorf("Failed to publish to user-mutation: %s", err)
			return nil, errors.Internal()
		}
	}

	return user, nil
}

func DeleteCurrentUser(ctx context.Context, req *commonModel.Empty) (*model.User, error) {
	userID, _ := ctx.Value(commonConstant.UserID).(string)
	return DeleteUser(ctx, &model.DeleteUserRequest{ID: userID})
}
