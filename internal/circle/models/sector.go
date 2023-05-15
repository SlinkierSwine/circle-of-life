package models

import (
	"circle-of-life/internal/core/db"

	"gorm.io/gorm"
)

type Sector struct {
    gorm.Model
    Name string `json:"name" gorm:"serializer:json"`
    Value float32 `json:"value" gorm:"serializer:json"`
    CircleID uint `json:"circle_id"`
    Circle Circle `json:"-"`
}


func (s *Sector) SaveSector() (*Sector, error) {

	var err error
	err = db.DB.Create(&s).Error
	if err != nil {
		return &Sector{}, err
	}
	return s, nil
}


func (s *Sector) ToRepresentation() map[string]interface{} {
    data := map[string]interface{} {
        "ID": s.Model.ID,
        "name": s.Name,
        "value": s.Value,
        "circle_id": s.CircleID,
    }
    return data
}
