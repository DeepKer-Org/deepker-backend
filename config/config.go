package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
)

var (
	DB         *gorm.DB
	DBUser     string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
)

// LoadConfig loads the database configuration and establishes the connection with PostgreSQL
func LoadConfig() {
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")

	if DBUser == "" || DBPassword == "" || DBName == "" || DBHost == "" || DBPort == "" {
		log.Fatal("Database configuration not set")
	}

	// Build the connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		DBHost, DBUser, DBPassword, DBName, DBPort)

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	log.Println("PostgreSQL database connected")
}

// CloseDB ensures the database connection is closed (if necessary)
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Println("Error getting database connection to close:", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	} else {
		log.Println("PostgreSQL connection closed")
	}
}
