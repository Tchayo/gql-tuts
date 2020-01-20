package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

type Author struct {
	gorm.Model
	Name  string `json:"name"`
	Email string	`gorm:"size:140;not null" ,json:"email"`
	Password  string    `gorm:"size:140;not null" json:"password"`
}

// HashPassword hash user password
func HashPassword(originalPassword string) (string, error)  {
	pass, err := bcrypt.GenerateFromPassword([]byte(originalPassword), 10)
	return string(pass), err
}

// ComparePassword compare password
func ComparePassword(password, hash string) bool {
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil {
		return true
	}
	return false
}

// Prepare : description
func (u *Author) Prepare() {
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.Password, _ = HashPassword(u.Password)
}

// SaveUser : description
func (u *Author) SaveUser(db *gorm.DB) (*Author, error) {

	u.Prepare()

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &Author{}, err
	}
	// hide user password
	u.Password = ""
	return u, nil
}