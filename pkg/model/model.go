package model

import (
	"time"
)

type Model struct {
	ID        int `json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

