package circle

import (
	"circle-of-life/internal/circle/models"
	"circle-of-life/internal/core/db"
	"errors"
	"fmt"

	"gorm.io/gorm"
)


func GetCircleByUsername(username string) (models.Circle, error) {
    var c models.Circle
    dbResult := db.DB.Joins("JOIN users on users.id = circles.user_id and users.username = ?", username).Find(&c)
	if dbResult.Error != nil || c.Model.ID == 0 {
		return c, errors.New(fmt.Sprintf("Circle not found!"))
	}
	return c, nil
}


func AppendSectorToCircle(username string, sector models.Sector) error {
    circle, err := GetCircleByUsername(username)
    if err != nil {
        return err
    }

    err = db.DB.Session(&gorm.Session{FullSaveAssociations: true}).Model(&circle).Association("Sectors").Append(&sector)
    if err != nil {
        return err
    }
	return nil
}


func GetSectorById(id int) (models.Sector, error) {
    var s models.Sector

	if err := db.DB.First(&s, id).Error; err != nil {
		return s, errors.New(fmt.Sprintf("Sector not found!"))
	}
	
	return s, nil
}

