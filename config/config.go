package config

import (
	"biometric-data-backend/utils"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	// Load app configuration first
	LoadAppConfig()

	// Use configuration from app config
	DBUser = App.DB.User
	DBPassword = App.DB.Password
	DBName = App.DB.Name
	DBHost = App.DB.Host
	DBPort = App.DB.Port

	// Build the connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, App.DB.SSLMode, App.DB.TimeZone)

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL: ", err)
	}
	log.Println("PostgreSQL database connected")

	// Run the migrations
	utils.ExecuteMigrations()
	// Load Redis configuration
	LoadRedisConfig()
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

var (
	RedisClient   *redis.Client
	RedisHost     string
	RedisPort     string
	RedisPassword string
)

// LoadRedisConfig initializes the Redis client with the configuration from app config.
func LoadRedisConfig() {
	// Check if Redis is enabled
	if !IsRedisEnabled() {
		log.Println("Redis is disabled. Skipping Redis initialization.")
		return
	}

	RedisHost = App.Redis.Host
	RedisPort = App.Redis.Port
	RedisPassword = App.Redis.Password

	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisHost + ":" + RedisPort,
		Password: RedisPassword,
		DB:       App.Redis.DB,
	})

	// Test the connection
	_, err := RedisClient.Ping(RedisClient.Context()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	log.Println("Redis connected successfully")
}

func CloseRedis() {
	if err := RedisClient.Close(); err != nil {
		log.Println("Error closing Redis connection:", err)
	} else {
		log.Println("Redis connection closed")
	}
}
