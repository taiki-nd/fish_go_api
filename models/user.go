package models

import "time"

type User struct {
	Id             uint      `json:"id" gorm:"primarykey"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Email          string    `json:"email"`
	Password       []byte    `json:"-"`
	PermissionType int64     `json:"permission_type"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
