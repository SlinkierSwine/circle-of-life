package user

import (
	"circle-of-life/internal/core/db"
	"errors"

	"golang.org/x/crypto/bcrypt"
)


func VerifyPassword(password,hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


func LoginCheck(username string, password string) (string,error) {
	
	var err error

	u := User{}

	err = db.DB.Model(User{}).Where("username = ?", username).Take(&u).Error

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


func GetUserByID(uid uint) (User, error) {

	var u User

	if err := db.DB.First(&u,uid).Error; err != nil {
		return u, errors.New("User not found!")
	}
	
	return u,nil

}
