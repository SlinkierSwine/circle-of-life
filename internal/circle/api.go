package circle

import (
	"circle-of-life/internal/circle/models"
	"circle-of-life/internal/core/db"
	"net/http"

	"github.com/gin-gonic/gin"
)


type SectorInput struct {
    Username string `json:"username" binding:"required"`
    Name string `json:"name" binding:"required"`
    Value float32 `json:"value"`
}


func CreateSector(c *gin.Context) {
    var input SectorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    
    var value float32
    if input.Value == 0 {
        value = 0
    } else {
        value = input.Value
    }

    sector := models.Sector{
        Name: input.Name,
        Value: value,
    }

    err := AppendSectorToCircle(input.Username, sector)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sector.ToRepresentation())
}


type UpdateSectorInput struct {
    Username string `json:"username" binding:"required"`
    Id int `json:"ID" binding:"required"`
    Name string `json:"name"`
    Value float32 `json:"value"`
}

func UpdateSector(c *gin.Context) {
    var input UpdateSectorInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    sector, err := GetSectorById(input.Id)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    sector.Name = input.Name   
    sector.Value = input.Value

    err = db.DB.Save(&sector).Error
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, sector.ToRepresentation())

}


type GetCircleInput struct {
	Username string `json:"username" binding:"required"`
}


func GetCircle(c *gin.Context) {
    var input GetCircleInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    username := input.Username
	circle, err := GetCircleByUsername(username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, circle.ToRepresentation())
}
