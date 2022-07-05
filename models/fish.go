package models

import (
	"time"
)

type Fish struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Fish      string    `json:"fish" gorm:"not null; size:256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
