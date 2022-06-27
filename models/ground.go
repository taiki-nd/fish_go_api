package models

import (
	"time"
)

type Ground struct {
	Id             uint            `json:"id" gorm:"primarykey"`
	Name           string          `json:"name" gorm:"not null; size:256"`
	Address        string          `json:"address" gorm:"not null; size:256"`
	Tell           string          `json:"tell" gorm:"not null; size:256"`
	Email          string          `json:"email" gorm:"not null; size:256"`
	Break          string          `json:"break" gorm:"not null; size:256"`
	Styles         []Style         `json:"styles" gorm:"many2many:ground_styles"`
	Howtos         []Howto         `json:"howtos" gorm:"many2many:ground_howtos"`
	Fishes         []Fish          `json:"fishes" gorm:"many2many:ground_fishes"`
	Price          string          `json:"price" gorm:"not null; size:256"`
	Url            string          `json:"url" gorm:"not null; size:256"`
	Feature        string          `json:"feature" gorm:"not null; size:256"`
	Rule           string          `json:"rule" gorm:"not null; size:256"`
	Other          string          `json:"other" gorm:"not null; size:256"`
	ImageUrl       string          `json:"image_url" gorm:"size:256"`
	GroundComments []GroundComment `json:"ground_comments" gorm:"foreignKey:GroundId"`
	CreatedAt      time.Time       `json:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"`
}
