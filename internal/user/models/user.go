package models

import (
	"circle-of-life/internal/circle/models"
	"circle-of-life/internal/core/db"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
    Circle models.Circle `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}


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


func (u *User) ToRepresentation() map[string]interface{} {
    data := map[string]interface{} {
        "id": u.Model.ID,
        "username": u.Username,
        "created_at": u.CreatedAt,
        "updated_at": u.UpdatedAt,
    }
    return data
}
