package config

import (
	"log"
	"os"
	"strconv"
	"time"
)

// AppConfig holds all application configuration
type AppConfig struct {
	Server ServerConfig
	DB     DatabaseConfig
	Redis  RedisConfig
	Cache  CacheConfig
	JWT    JWTConfig
	CORS   CORSConfig
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Port string
	Host string
	Mode string // gin mode: debug, release, test
}

// DatabaseConfig holds database configuration
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
	TimeZone string
}

// RedisConfig holds Redis configuration
type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	Enabled  bool
}

// CacheConfig holds cache-related configuration
type CacheConfig struct {
	Enabled    bool
	DefaultTTL time.Duration
}

// JWTConfig holds JWT configuration
type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

// CORSConfig holds CORS configuration
type CORSConfig struct {
	AllowedOrigin string
}

var App *AppConfig

// LoadAppConfig loads all application configuration from environment variables
func LoadAppConfig() *AppConfig {
	App = &AppConfig{
		Server: ServerConfig{
			Port: getEnvOrDefault("SERVER_PORT", "8080"),
			Host: getEnvOrDefault("SERVER_HOST", "0.0.0.0"),
			Mode: getEnvOrDefault("GIN_MODE", "debug"),
		},
		DB: DatabaseConfig{
			Host:     getEnvOrDefault("DB_HOST", "localhost"),
			Port:     getEnvOrDefault("DB_PORT", "5432"),
			User:     getEnvOrDefault("DB_USER", "postgres"),
			Password: getEnvOrDefault("DB_PASSWORD", ""),
			Name:     getEnvOrDefault("DB_NAME", "deepker"),
			SSLMode:  getEnvOrDefault("SSL_MODE", "disable"),
			TimeZone: getEnvOrDefault("TIME_ZONE", "UTC"),
		},
		Redis: RedisConfig{
			Host:     getEnvOrDefault("REDIS_HOST", "localhost"),
			Port:     getEnvOrDefault("REDIS_PORT", "6379"),
			Password: getEnvOrDefault("REDIS_PASSWORD", ""),
			DB:       getEnvAsInt("REDIS_DB", 0),
			Enabled:  getEnvAsBool("REDIS_ENABLED", true),
		},
		Cache: CacheConfig{
			Enabled:    getEnvAsBool("CACHE_ENABLED", true),
			DefaultTTL: getEnvAsDuration("CACHE_DEFAULT_TTL", 5*time.Minute),
		},
		JWT: JWTConfig{
			SecretKey: getEnvOrDefault("JWT_SECRET_KEY", ""),
			ExpiresIn: getEnvAsDuration("JWT_EXPIRES_IN", 24*time.Hour),
		},
		CORS: CORSConfig{
			AllowedOrigin: getEnvOrDefault("ALLOWED_ORIGIN", "http://localhost:3000"),
		},
	}

	// Validate required configuration
	if App.DB.Password == "" {
		log.Fatal("DB_PASSWORD is required")
	}
	if App.JWT.SecretKey == "" {
		log.Fatal("JWT_SECRET_KEY is required")
	}

	log.Printf("Configuration loaded: Cache enabled: %v, Redis enabled: %v", 
		App.Cache.Enabled, App.Redis.Enabled)

	return App
}

// Helper functions for environment variable parsing
func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
		log.Printf("Warning: Invalid boolean value for %s: %s, using default: %v", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
		log.Printf("Warning: Invalid integer value for %s: %s, using default: %d", key, value, defaultValue)
	}
	return defaultValue
}

func getEnvAsDuration(key string, defaultValue time.Duration) time.Duration {
	if value := os.Getenv(key); value != "" {
		if parsed, err := time.ParseDuration(value); err == nil {
			return parsed
		}
		log.Printf("Warning: Invalid duration value for %s: %s, using default: %v", key, value, defaultValue)
	}
	return defaultValue
}

// IsCacheEnabled returns whether caching is enabled
func IsCacheEnabled() bool {
	return App != nil && App.Cache.Enabled && App.Redis.Enabled
}

// IsRedisEnabled returns whether Redis is enabled
func IsRedisEnabled() bool {
	return App != nil && App.Redis.Enabled
}