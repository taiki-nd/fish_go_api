package models

import (
	"time"
)

type PostComment struct {
	Id        uint      `json:"id" gorm:"primarykey"`
	Name      string    `json:"name" gorm:"not null; size:256; default:gest"`
	Content   string    `json:"content" gorm:"not null; size:256"`
	ImageUrl  string    `json:"image_url" gorm:"size:256"`
	Filename  string    `json:"filename" gorm:"size:256"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
