package controller

import (
	"context"
	e "errors"

	commonConstant "github.com/emitra-labs/common/constant"
	"github.com/emitra-labs/common/errors"
	"github.com/emitra-labs/common/id"
	"github.com/emitra-labs/common/log"
	commonModel "github.com/emitra-labs/common/model"
	"github.com/emitra-labs/common/paginator"
	"github.com/emitra-labs/common/validator"
	"github.com/emitra-labs/identity-service/db"
	"github.com/emitra-labs/identity-service/model"
	"gorm.io/gorm"
)

func CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.Session, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	session := &model.Session{
		ID:        id.New(),
		Token:     id.New(36),
		UserID:    req.UserID,
		ExpiresAt: req.ExpiresAt,
	}

	if err := db.DB.WithContext(ctx).Create(session).Error; err != nil {
		log.Errorf("Failed to create session: %s", err)
		return nil, errors.Internal()
	}

	return session, nil
}

func GetSessions(ctx context.Context, req *model.GetSessionsRequest) (*model.GetSessionsResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).
		Model(&model.Session{}).
		Where("user_id = ?", req.UserID)

	data, pagination, err := paginator.Paginate[model.Session](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetSessionsResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}

func GetCurrentUserSessions(ctx context.Context, req *model.GetCurrentUserSessionsRequest) (*model.GetSessionsResponse, error) {
	userID, _ := ctx.Value(commonConstant.UserID).(string)

	return GetSessions(ctx, &model.GetSessionsRequest{
		PaginationRequest: req.PaginationRequest,
		UserID:            userID,
	})
}

func GetSession(ctx context.Context, req *model.GetSessionRequest) (*model.Session, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	session := new(model.Session)

	tx := db.DB.WithContext(ctx).Where("expires_at > NOW()")

	if req.ID != "" {
		tx = tx.Where("id = ?", req.ID)
	}

	if req.Token != "" {
		tx = tx.Where("token = ?", req.Token)
	}

	if req.UserID != "" {
		tx = tx.Where("user_id = ?", req.UserID)
	}

	err := tx.Take(session).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound("Session not found")
		}

		log.Errorf("Failed to view session: %s", err)
		return nil, errors.Internal()
	}

	return session, nil
}

func GetCurrentUserSession(ctx context.Context, req *model.GetCurrentUserSession) (*model.Session, error) {
	userID, _ := ctx.Value(commonConstant.UserID).(string)

	return GetSession(ctx, &model.GetSessionRequest{
		ID:     req.ID,
		UserID: userID,
	})
}

func GetCurrentUserCurrentSession(ctx context.Context, req *commonModel.Empty) (*model.Session, error) {
	sessionID, _ := ctx.Value(commonConstant.SessionID).(string)
	userID, _ := ctx.Value(commonConstant.UserID).(string)

	return GetSession(ctx, &model.GetSessionRequest{
		ID:     sessionID,
		UserID: userID,
	})
}

func DeleteSession(ctx context.Context, req *model.DeleteSessionRequest) (*model.Session, error) {
	s, err := GetSession(ctx, &model.GetSessionRequest{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := db.DB.Delete(s).Error; err != nil {
		log.Errorf("Failed to delete session: %s", err)
		return nil, errors.Internal()
	}

	return s, nil
}

func DeleteOldSessions(ctx context.Context, req *model.DeleteOldSessionsRequest) (*commonModel.Empty, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	// Delete old sessions, preserve 3 newest
	sql := `DELETE FROM sessions WHERE user_id = ? AND id NOT IN (
		SELECT id FROM sessions WHERE user_id = ? ORDER BY created_at DESC LIMIT 3
	)`
	if err := db.DB.Exec(sql, req.UserID, req.UserID).Error; err != nil {
		log.Errorf("Failed to delete old sessions: %s", err)
		return nil, errors.Internal()
	}

	return &commonModel.Empty{}, nil
}
