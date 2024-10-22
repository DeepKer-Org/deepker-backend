package service

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("mySecretKey")

// GenerateToken genera un token JWT con roles
func GenerateToken(username string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"roles":    roles,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // 24 hrs to expire
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
