package model

type SignUpRequest struct {
	Name     string `json:"name" validate:"required,lte=30" example:"User"`
	Email    string `json:"email" validate:"required,email,lte=30" example:"user@example.com"`
	Password string `json:"password" validate:"required,gte=8,lte=64" example:"SuperSecret"`
}
