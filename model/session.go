package model

import "time"

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

type DeleteOldSessionsRequest struct {
	UserID string `validate:"required"`
}

type GetSessionRequest struct {
	ID     string `params:"id" path:"id" validate:"required_without=Token,required_without=UserID"`
	Token  string `validate:"required_without=ID"`
	UserID string `validate:"required_without=ID"`
}
