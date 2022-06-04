package models

import (
	"time"
)

type Howto struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Howto     string    `json:"howto" gorm:"not null; size:256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
