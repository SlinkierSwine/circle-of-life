package user

import (
	"circle-of-life/internal/circle"

	"gorm.io/gorm"
)


type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"password"`
    Circle circle.Circle
}
