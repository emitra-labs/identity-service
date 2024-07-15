package model

import (
	"time"

	commonModel "github.com/ukasyah-dev/common/model"
)

type User struct {
	ID         string    `gorm:"primaryKey" json:"id" example:"h8WpMrLeTA7mgyDGCtEkiX"`
	Name       string    `json:"name" example:"User"`
	Email      string    `gorm:"unique" json:"email" example:"user@example.com"`
	Password   string    `json:"-"`
	Status     string    `json:"status" example:"active"`
	SuperAdmin bool      `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type CreateUserRequest struct {
	Name       string `json:"name" validate:"required,lte=30" example:"User"`
	Email      string `json:"email" validate:"required,email,lte=30" example:"user@example.com"`
	Password   string `json:"password" validate:"required,gte=8,lte=64" example:"SuperSecret"`
	Status     string `json:"status" example:"active"`
	SuperAdmin bool   `json:"-"`
}

type GetUsersRequest struct {
	commonModel.PaginationRequest
}

type GetUsersResponse struct {
	commonModel.PaginationResponse
	Data []*User `json:"data"`
}

type GetUserRequest struct {
	ID     string `params:"id" path:"id" validate:"required_without=Email"`
	Email  string `validate:"required_without=ID"`
	Status string
}
