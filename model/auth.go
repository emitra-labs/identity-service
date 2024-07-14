package model

import "time"

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,lte=30" example:"User"`
	Email    string `json:"email" validate:"required,email,lte=30" example:"user@example.com"`
	Password string `json:"password" validate:"required,gte=8,lte=64" example:"SuperSecret"`
}

type SignInRequest struct {
	Email    string `json:"email" validate:"required,email" example:"user@example.com"`
	Password string `json:"password" validate:"required" example:"SuperSecret"`
}

type Access struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type Refresh struct {
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expiresAt"`
}

type SignInResponse struct {
	Access  Access  `json:"access"`
	Refresh Refresh `json:"refresh"`
}
