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
