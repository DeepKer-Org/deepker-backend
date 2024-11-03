package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// GenerateToken generates a JWT token
func GenerateToken(email string, roles []string, additionalClaims map[string]interface{}) (string, error) {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		return "", fmt.Errorf("JWT_SECRET_KEY not set in environment")
	}
	claims := jwt.MapClaims{
		"email": email,
		"roles": roles,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 hrs to expire
	}

	for key, value := range additionalClaims {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
