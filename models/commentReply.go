package models

import (
	"time"
)

type CommentReply struct {
	Id              uint      `json:"id" gorm:"primarykey"`
	Name            string    `json:"name" gorm:"not null; size:256"`
	Content         string    `json:"content" gorm:"not null; size:256"`
	GroundCommentId uint      `json:"ground_comment_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
