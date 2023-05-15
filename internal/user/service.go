package user

import (
	"circle-of-life/internal/user/models"
	"circle-of-life/internal/core/db"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)


func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func LoginCheck(username string, password string) (string,error) {
	
	var err error

	u := models.User{}

	err = db.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
	
}


func GetUserByID(uid uint) (models.User, error) {

	var u models.User

	if err := db.DB.First(&u,uid).Error; err != nil {
		return u, errors.New(fmt.Sprintf("User not found! %s", err.Error()))
	}
	
	return u, nil

}

func GetUserByUsername(username string) (models.User, error) {
    var u models.User

	if err := db.DB.Where("username = ?", username).First(&u).Error; err != nil {
		return u, errors.New(fmt.Sprintf("User not found! %s", err.Error()))
	}
	
	return u, nil
}
