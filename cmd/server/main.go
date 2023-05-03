package main

import (
	"circle-of-life/internal/core/db"
	"circle-of-life/internal/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()

    err := db.DB.AutoMigrate(user.User{})
    if err != nil {
        log.Fatal(err.Error())
    }

    r := gin.Default()

    api := r.Group("/api")

    api.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
    })

    api.POST("/register", user.Register)

    r.Run()
}
