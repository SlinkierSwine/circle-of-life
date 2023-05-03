package user

import (
	"html"
	"strings"

	"circle-of-life/internal/core/db"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func (u *User) SaveUser() (*User, error) {

	var err error
	err = db.DB.Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password),bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username 
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
