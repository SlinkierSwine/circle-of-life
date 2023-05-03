package db

import (
	"fmt"
	"log"
	"os"

    "gorm.io/gorm"
    "gorm.io/driver/postgres"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDB(){

	err := godotenv.Load(".env")

	if err != nil {
	  log.Fatalf("Error loading .env file")
	}	
	
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USERNAME")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

    DBURL := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName, DbPort)
	
    DB, err = gorm.Open(postgres.Open(DBURL), &gorm.Config{})

	if err != nil {
        log.Fatal("Db connection error:", err)
		panic("Cannot connect to database")
	} else {
		fmt.Println("Connection to database is established")
	}
		
}
