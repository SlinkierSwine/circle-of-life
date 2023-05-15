package user

import (
	circleModels "circle-of-life/internal/circle/models"
	"circle-of-life/internal/core/db"
	"circle-of-life/internal/user/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Register(c *gin.Context){
    var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    u := models.User{
        Username: input.Username,
        Password: input.Password,
        Circle: circleModels.Circle{
            Sectors: []circleModels.Sector{},
        },
    }

    err := db.DB.Create(&u).Error

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
    db.DB.Save(&u)

	c.JSON(http.StatusCreated, gin.H{"message": "registration success"})

}


type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func Login(c *gin.Context) {
	
	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	token, err := LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})

}


func CurrentUser(c *gin.Context){

	user_id, err := ExtractTokenID(c)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u, err := GetUserByID(user_id)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, u.ToRepresentation())
}
