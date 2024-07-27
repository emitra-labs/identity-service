package model

import (
	"time"

	commonModel "github.com/emitra-labs/common/model"
)

type Session struct {
	ID        string    `gorm:"primaryKey" json:"id" example:"iksMPzcgzo4BvRiDycM74L"`
	Token     string    `gorm:"index:idx_user_session_token,unique" json:"-"`
	UserID    string    `gorm:"index:idx_user_session_token,unique" json:"userId" example:"h8WpMrLeTA7mgyDGCtEkiX"`
	User      *User     `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type CreateSessionRequest struct {
	UserID    string    `validate:"required"`
	ExpiresAt time.Time `validate:"required"`
}

type GetSessionsRequest struct {
	commonModel.PaginationRequest
	UserID string `validate:"required"`
}

type GetCurrentUserSessionsRequest struct {
	commonModel.PaginationRequest
}

type GetSessionsResponse struct {
	commonModel.PaginationResponse
	Data []*Session `json:"data"`
}

type GetSessionRequest struct {
	ID     string `validate:"required_without=Token,required_without=UserID"`
	Token  string `validate:"required_without=ID"`
	UserID string `validate:"required_without=ID"`
}

type GetCurrentUserSession struct {
	ID string `params:"sessionId" path:"sessionId"`
}

type DeleteSessionRequest struct {
	ID string `validate:"required"`
}

type DeleteOldSessionsRequest struct {
	UserID string `validate:"required"`
}
