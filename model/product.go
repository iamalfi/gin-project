package model

import (
	"github.com/google/uuid"
)

type Product struct {
	Base
	UserID      uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Title       string    `gorm:"not null" json:"title"`
	Image       string    `json:"image"`
	Color       string    `json:"color"`
	Size        string    `json:"size"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
}
