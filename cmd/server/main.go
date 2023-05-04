package main

import (
	"circle-of-life/internal/circle"
	"circle-of-life/internal/core/db"
	"circle-of-life/internal/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db.ConnectDB()

    err := db.DB.AutoMigrate(user.User{}, circle.Circle{}, circle.Sector{})
    if err != nil {
        log.Fatal(err.Error())
    }

    r := gin.Default()

    public := r.Group("/api")

    public.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
    })

    // Auth
    public.POST("/register", user.Register)
    public.POST("/login", user.Login)

    protected := r.Group("/api")

    // User
    protected.Use(user.JwtAuthMiddleware())
    protected.GET("/me", user.CurrentUser)

    r.Run()
}
