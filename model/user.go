package model

import "time"

type User struct {
	ID         string    `gorm:"primaryKey" json:"id" example:"h8WpMrLeTA7mgyDGCtEkiXQ"`
	Email      string    `gorm:"unique" json:"email" example:"user@example.com"`
	Name       string    `json:"name" example:"User"`
	Password   string    `json:"-"`
	Status     string    `json:"status" example:"active"`
	SuperAdmin bool      `json:"-"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
