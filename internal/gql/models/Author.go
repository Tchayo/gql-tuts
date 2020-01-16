package models

import "github.com/jinzhu/gorm"

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string	`gorm:"size:140;not null" ,json:"email"`
}