package database

import (
	"fmt"
	"log"
	"os"

	"github.com/CodeWithPreet/fiber-api/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type DBIntance struct {
	Db *gorm.DB
}

var DB DBIntance

func ConnectDB()  {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Read database credentials from environment variables
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	// PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	db ,err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err!= nil{
		log.Fatal("Failed to connect \n", err.Error()) 
		os.Exit(2)
	}
	log.Println("Successfully Connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DBMigation(db)

	DB= DBIntance{Db: db}

	
}

func DBMigation( db *gorm.DB)  {

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Order{})
	
}