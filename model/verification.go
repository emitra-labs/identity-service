package model

import "time"

type Verification struct {
	ID        string    `gorm:"primaryKey" json:"id" example:"iksMPzcgzo4BvRiDycM74L"`
	Token     string    `gorm:"index:idx_user_verification_token,unique" json:"-"`
	UserID    string    `gorm:"index:idx_user_verification_token,unique" json:"userId" example:"h8WpMrLeTA7mgyDGCtEkiX"`
	User      *User     `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	Status    string    `json:"status" example:"pending"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type CreateVerificationRequest struct {
	UserID    string    `validate:"required"`
	ExpiresAt time.Time `validate:"required"`
}

type VerifyRequest struct {
	UserID string `json:"userId" validate:"required"`
	Token  string `json:"token" validate:"required"`
}
