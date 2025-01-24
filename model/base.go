package model

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid; unique; primary_key; column:id; default:uuid_generate_v4()"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:createdAt"`
}
