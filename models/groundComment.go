package models

import (
	"time"
)

type GroundComment struct {
	Id             uint           `json:"id" gorm:"primarykey"`
	Name           string         `json:"name" gorm:"not null; size:256"`
	Content        string         `json:"content" gorm:"not null; size:256"`
	Rate           int64          `json:"Rate"`
	Size           int64          `json:"Size"`
	Nice           int64          `json:"Nice" gorm:"default:0"`
	CommentReplies []CommentReply `json:"comment_replies" gorm:"foreignKey:GroundCommentId"`
	GroundId       uint           `json:"ground_id"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}
