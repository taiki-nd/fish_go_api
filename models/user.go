package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id              uint      `json:"id" gorm:"primarykey"`
	FirstName       string    `json:"first_name" gorm:"not null; size:256"`
	LastName        string    `json:"last_name" gorm:"not null; size:256"`
	Email           string    `json:"email" gorm:"unique; not null; size:256"`
	Password        []byte    `json:"-" gorm:"not null;"`
	PasswordConfirm []byte    `json:"-"`
	PermissionType  string    `json:"permission_type" gorm:"not null; size:256"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (user *User) SetPassword(password string) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedPassword
	user.PasswordConfirm = []byte("--------")
}
