package models

import (
	"time"
)

type Style struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Style     string    `json:"style" gorm:"not null; size:256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
