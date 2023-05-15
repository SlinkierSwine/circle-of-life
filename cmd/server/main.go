package main

import (
	"circle-of-life/internal/circle"
	circleModels "circle-of-life/internal/circle/models"
	"circle-of-life/internal/core/db"
	"circle-of-life/internal/user"
	userModels "circle-of-life/internal/user/models"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db.ConnectDB()

    err := db.DB.AutoMigrate(userModels.User{}, circleModels.Circle{}, circleModels.Sector{})
    if err != nil {
        log.Fatal(err.Error())
    }
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatalf("unable to load .env file")
    }

    if ginMode, exists := os.LookupEnv("GIN_MODE"); exists && ginMode == "release" {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.Default()

    public := r.Group("/api")

    public.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
    })

    // Auth
    public.POST("/register", user.Register)
    public.POST("/login", user.Login)

    public.POST("/circle", circle.GetCircle)
    public.POST("/sector", circle.CreateSector)
    public.PUT("/sector", circle.UpdateSector)

    protected := r.Group("/api")

    // User
    protected.Use(user.JwtAuthMiddleware())
    protected.GET("/me", user.CurrentUser)

    port, exists := os.LookupEnv("PORT")
    if !exists {
        port = "8000"
    }

    r.Run(fmt.Sprintf(":%s", port))
}
