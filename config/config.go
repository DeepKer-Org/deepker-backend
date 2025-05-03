package config

import (
	"biometric-data-backend/utils"
	"fmt"
	"log"
	"os"

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
	DBUser = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName = os.Getenv("DB_NAME")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSL_MODE")
	TimeZone := os.Getenv("TIME_ZONE")

	if DBUser == "" || DBPassword == "" || DBName == "" || DBHost == "" || DBPort == "" {
		log.Fatal("Database configuration not set")
	}

	// Build the connection string for PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort, sslmode, TimeZone)

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

// LoadRedisConfig initializes the Redis client with the configuration from environment variables.
func LoadRedisConfig() {
	// Check if caching is enabled
    cacheEnabled := os.Getenv("CACHE_ENABLED") != "false"
    if !cacheEnabled {
        log.Println("Cache is disabled. Skipping Redis initialization.")
        return
    }

	RedisHost = os.Getenv("REDIS_HOST")
	RedisPort = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")

	if RedisHost == "" || RedisPort == "" {
		log.Fatal("Redis configuration not set")
	}

	// Initialize Redis client
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     RedisHost + ":" + RedisPort,
		Password: RedisPassword, // Leave empty if no password
		DB:       0,             // Use default DB
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
