package storage

import (
	"fmt"
	"log"
	"os"

	"github.com/ivanmartinezmorales/life-server/server/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB The global of our DB
var DB *gorm.DB

// ConnectDB connects to the db and runs automigrations
func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file \n", err)
	}

	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))

	log.Print("Connecting to pSQL DB")

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
		os.Exit(2)
	}

	log.Println("Connected to db")

	DB.Logger = logger.Default.LogMode(logger.Info)

	log.Print("Running migrations...")

	DB.AutoMigrate(&models.User{}, &models.Claims{})
}
