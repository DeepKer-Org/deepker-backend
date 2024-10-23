package service

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// GenerateToken generates a JWT token
func GenerateToken(email string, roles []string, additionalClaims map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"roles": roles,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // 24 hrs to expire
	}

	for key, value := range additionalClaims {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
