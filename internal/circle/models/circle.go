package models

import (
	"circle-of-life/internal/core/db"

	"gorm.io/gorm"
)


type Circle struct {
    gorm.Model
    Sectors []Sector `json:"sectors" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
    UserID int `json:"user_id"`
}


func (c *Circle) SaveCircle() (*Circle, error) {

	var err error
	err = db.DB.Create(&c).Error
	if err != nil {
		return &Circle{}, err
	}
	return c, nil
}


func (c *Circle) ToRepresentation() map[string]interface{} {
    sectors := []Sector{}
    db.DB.Preload("Circle").Find(&sectors)
    data := map[string]interface{} {
        "ID": c.Model.ID,
        "sectors": sectors,
    }
    return data
}
